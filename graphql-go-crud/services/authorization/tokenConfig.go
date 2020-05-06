package authorization

import (
	"os"
	"time"
	"log"
	b64 "encoding/base64"
	"github.com/dgrijalva/jwt-go"

	"graphql-go-crud/config"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func  GenerateToken(username,password string)  (string,bool) {
	secretkey, _ := b64.StdEncoding.DecodeString(os.Getenv("secretkey"))
	serverUser, _ := b64.StdEncoding.DecodeString(os.Getenv("serverUser"))
	serverPassword, _ :=  b64.StdEncoding.DecodeString(os.Getenv("serverPassword"))
	user, _ := b64.StdEncoding.DecodeString(username)
	pass,_ := b64.StdEncoding.DecodeString(password)

	if string(serverPassword) != string(pass) || string(serverUser) != string(user) {
		return "UnAuthorized",false
	}
	expirationTime := time.Now().Add(time.Duration(config.Config.TokenExpiry) * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		return "Internal Server Error",false
	}
	return tokenString,true
	
}

func ValidateToken(token string) (string,bool) {
	claims := &Claims{}
	jwtKey,_ := b64.StdEncoding.DecodeString(os.Getenv("secretkey"))
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Print(err)
			return "UnAuthorized or Invalid token",false
		}
		log.Print(err)
		return "UnAuthorized or Invalid token",false
	}
	if !tkn.Valid {
		return "Invalid Token",false
	}
	return "",true

}
