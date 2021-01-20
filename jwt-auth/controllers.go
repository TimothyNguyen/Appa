package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func login(c *gin.Context) {
	var u User      // request form
	var f Feedback  // feedback to client
	var result User // result from database
	var err error
	if err = c.ShouldBind(&u); err != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Invalid login form")
		c.JSON(http.StatusUnprocessableEntity, f)
		return
	}
	// TODO: check if the user is in the cache!

	// Find User by email (done)
	filter := bson.M{"email": u.Email}
	err = userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "User not found")
		c.JSON(http.StatusUnauthorized, f)
		return
	}

	// Authenticate user password (done)
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(u.Password))
	if err != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Email or password incorrect")
		c.JSON(http.StatusUnauthorized, f)
		return
	}

	// Create token pari
	td, err := createToken(result.ID)
	if err != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Token not created")
		f.Msgs = append(f.Msgs, err.Error())
		c.JSON(http.StatusUnprocessableEntity, f)
		return
	}

	// save token into redis (done)
	saveErr := createAuth(result.ID, td)
	if saveErr != nil {
		f.Status = "unsuccess"
		f.Msgs = append(f.Msgs, "Token created but not saved")
		f.Msgs = append(f.Msgs, saveErr.Error())
		c.JSON(http.StatusUnprocessableEntity, f)
		return
	}

	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
		"token":         td.RefreshToken,
	}
	f.Status = "success"
	f.Data = tokens
	c.JSON(http.StatusOK, f)
}

// TODO: implement register router
func register(c *gin.Context) {
	var u User      // request form
	var f Feedback  // feedback to client
	var result User // result from database

	if err := c.ShouldBind(&u); err != nil {
		f.Status = "400"
		f.Msgs = append(f.Msgs, "Invalid register form")
		c.JSON(http.StatusUnprocessableEntity, f)
		return
	}
	// TODO: check if the user is in the cache!

	// Find User by email (done)
	filter := bson.M{"email": u.Email}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		f.Status = "400"
		f.Msgs = append(f.Msgs, "Email exists")
		c.JSON(http.StatusUnauthorized, f)
		return
	}

	// TODO: check email and password format is valid
	validFormat, message := isEmailValid(u.Email, u.Password)
	if !validFormat {
		f.Status = "400"
		f.Msgs = append(f.Msgs, message)
		c.JSON(http.StatusUnauthorized, f)
		return
	}

	// Hash User password and store it into database
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	/*
		insertUser := User{
			Id:
			Name:                 u.Name,
			Email:                u.Email,
			Password:             string(bytes),
			Date:                 primitive.Timestamp{T: uint32(time.Now().Unix())},
			VerificationURLCode:  ,
			PasswordResetURLCode: "",
		}
	*/
	insertUser := bson.M{
		"name":                    u.Name,
		"email":                   u.Email,
		"password":                string(bytes),
		"date":                    primitive.Timestamp{T: uint32(time.Now().Unix())},
		"verification_url_code":   uuid,
		"password_reset_url_code": "",
	}

	_, err = userCollection.InsertOne(context.TODO(), insertUser)
	if err != nil {
		panic(err)
	}
	f.Status = "success"
	c.JSON(http.StatusOK, insertUser)
}

func refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	tokenString := mapToken["refresh_token"]
	token, err := VerifyToken(tokenString, refreshSecret)
	fmt.Println(token)

	if err != nil {
		f := xFeedback("Refresh token not valid")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	// get id from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		f := xFeedback("Refresh token not valid")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	uuid := claims["uuid"].(string)
	userIDString, err := rdb.Get(uuid).Result()
	userIDHex, err := hex.DecodeString(userIDString)
	if err != nil || len(userIDHex) != 12 {
		f := xFeedback("_id not valid")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	var userID [12]byte
	copy(userID[:], userIDHex[:])
	_, err = rdb.Del(uuid).Result()
	if err != nil {
		f := xFeedback("Refresh token not deleted")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	td, err := createToken(userID)
	if err != nil {
		f := xFeedback("Fail to generate new token")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	saveErr := createAuth(userID, td)
	if saveErr != nil {
		f := xFeedback("Fail to save new token")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}
	f := yFeedback(tokens)
	c.JSON(http.StatusOK, f)
}

func logout(c *gin.Context) {
	bearToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		f := xFeedback("No token found in the header")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	tokenString := strArr[1]
	token, err := VerifyToken(tokenString, accessSecret)
	if err != nil {
		f := xFeedback("Access token not valid")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	// get id from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		f := xFeedback("Access token not valid")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	uuid := claims["uuid"].(string)
	_, err = rdb.Del(uuid).Result()
	if err != nil {
		f := xFeedback("token not found")
		c.JSON(http.StatusUnauthorized, f)
		return
	}
	f := yFeedback(nil)
	c.JSON(http.StatusOK, f)
	return
}

// isEmailValid check if the email provided passes the required
// structure and length test. It also checks the domain has a valid
// MX record
func isEmailValid(userEmail string, userPassword string) (bool, string) {
	// 1. Looking at email
	if len(userEmail) < 8 || len(userEmail) > 254 {
		return false, "The length of the email isn't valid."
	}
	if !emailRegex.MatchString(userEmail) {
		return false, "The provided email isn't a valid email"
	}
	parts := strings.Split(userEmail, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false, "The provided email isn't a valid email"
	}

	// 2. Looking at password
	if len(userPassword) < 10 || len(userPassword) > 100 {
		return false, "The length of the password is invalid. Needs to be between 10 and 100 characters."
	}

	return true, "The email and password is valid"
}
