package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

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

	// Hash User password and store it into database
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
	insertUser := bson.M{
		"name":     u.Name,
		"email":    u.Email,
		"password": string(bytes),
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
