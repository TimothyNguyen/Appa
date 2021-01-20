package main

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type signingMethod = jwt.SigningMethodHMAC

var accessSecret = []byte(os.Getenv("JWT_ACCESS_KEY")) // change this to os.Getenv("ACCESS_SECRET") for deployment

var refreshSecret = []byte(os.Getenv("JWT_REFRESH_KEY")) // change this to os.Getenv("REFRESH_SECRET") for deployment

var dbName = "appa-test" // change this to "uplink" for deployment

var collectionName = "users"

var mongodbURL = os.Getenv("MONGODB_URL")

var redisURL = "redis:6379"
