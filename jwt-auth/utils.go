package main // store token into redis
import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func xFeedback(msg string) Feedback {
	var f Feedback
	f.Status = "unsuccess"
	f.Msgs = append(f.Msgs, msg)
	return f
}

func yFeedback(data interface{}) Feedback {
	var f Feedback
	f.Status = "success"
	f.Data = data
	return f
}

// store token into redis
func createAuth(userID primitive.ObjectID, td UserToken) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := rdb.Set(td.AccessUUID, userID.Hex(), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := rdb.Set(td.RefreshUUID, userID.Hex(), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

// TokenAuthMiddleware check the access token is valid before execute next router
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		strArr := strings.Split(bearToken, " ")
		if len(strArr) != 2 {
			f := xFeedback("No token found in the header")
			c.JSON(http.StatusUnauthorized, f)
			return
		}
		tokenString := strArr[1]
		_, err := VerifyToken(tokenString, accessSecret)
		if err != nil {
			f := xFeedback("Access token not valid")
			c.JSON(http.StatusUnauthorized, f)
			return
		}
		c.Next()
	}
}

// VerifyToken verifies the token is valid using a given key
func VerifyToken(tokenString string, secret []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*signingMethod); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func createTokenWithUser(id primitive.ObjectID, result User) (UserToken, error) {
	var err error
	td := UserToken{}
	empty := UserToken{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()
	td.AccessUUID = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 14).Unix()
	td.RefreshUUID = uuid.NewV4().String()
	td.User = result
	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true
	atClaim["access_uuid"] = td.AccessUUID
	atClaim["_id"] = id
	atClaim["expire"] = td.AtExpires
	atClaim["user"] = td.User
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	td.AccessToken, err = at.SignedString(accessSecret)
	if err != nil {
		return empty, err
	}
	rtClaim := jwt.MapClaims{}
	rtClaim["refresh_uuid"] = td.RefreshUUID
	rtClaim["_id"] = id
	rtClaim["expire"] = td.RtExpires
	rtClaim["user"] = td.User
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaim)
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return empty, err
	}
	return td, nil
}

func createToken(id primitive.ObjectID) (UserToken, error) {
	var err error
	td := UserToken{}
	empty := UserToken{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()
	td.AccessUUID = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 14).Unix()
	td.RefreshUUID = uuid.NewV4().String()
	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true
	atClaim["access_uuid"] = td.AccessUUID
	atClaim["_id"] = id
	atClaim["expire"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	td.AccessToken, err = at.SignedString(accessSecret)
	if err != nil {
		return empty, err
	}
	rtClaim := jwt.MapClaims{}
	rtClaim["refresh_uuid"] = td.RefreshUUID
	rtClaim["_id"] = id
	rtClaim["expire"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaim)
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return empty, err
	}
	return td, nil
}

/*
"id":              result.ID.String(),
		"name":            result.Name,
		"email":           result.Email,
		"phone_number":    result.Password,
		"github_username": result.GithubUsername,
		"linkedin":        result.Linkedin,
		"user_type":       result.UserType,
*/
