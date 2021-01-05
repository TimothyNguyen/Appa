package main

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type signingMethod = jwt.SigningMethodHMAC

var accessSecret = []byte(os.Getenv("APPA_DEV_JWT_ACCESS_SECRET")) // change this to os.Getenv("ACCESS_SECRET") for deployment

var refreshSecret = []byte(os.Getenv("APPA_DEV_JWT_REFRESH_SECRET")) // change this to os.Getenv("REFRESH_SECRET") for deployment

var collectionName = "users"

var mongodbURL = os.Getenv("APPA_DEV_MONGODB_URL")

var redisURL = "redis:6379"
