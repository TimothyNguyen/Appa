package main // store token into redis
import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
