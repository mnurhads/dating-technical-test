package Auth

import (
	"errors"
	"time"
	"os"
	"github.com/joho/godotenv"
	"path/filepath"
	"log"
	"github.com/dgrijalva/jwt-go"
)

type JWTClaim struct { 
	ClientKey 		string
	ClientSecret	string
	jwt.StandardClaims
}

func goDotEnvVariable(key string) string {

	// load .env file
  path := os.Getenv("GIN_ENV")
  folder := ""
  if path == "production" {
		folder = "/var/www/go/connector-permata/"
	}
  //fmt.Println(filepath.Join(path,folder, ".env"))

  err := godotenv.Load(filepath.Join(folder, ".env"))

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}


var	jwtKey = []byte(goDotEnvVariable("JWT_KEY"))

func GenerateJWT(clientKey string, clientsecret string) (tokenString string, err error) {
	// if jwtKey == "" {
	// 	jwtKey = []byte("secretnominapay")
	// }
	expirationTime := time.Now().Add(time.Minute * 15)
	claims:= &JWTClaim{
		ClientKey: clientKey,
		ClientSecret: clientsecret,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

// func GetEmail(signedToken string) (email string, err error) {
// 	token, err := jwt.ParseWithClaims(
// 		signedToken,
// 		&JWTClaim{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(jwtKey), nil
// 		},
// 	)
// 	if err != nil {
// 		return
// 	}

// 	claims, ok := token.Claims.(*JWTClaim)
// 	if !ok {
// 		err = errors.New("couldn't parse claims")
// 		return
// 	}
// 	email = claims.Email
// 	return
// }
