package core

import (
	"fmt"
	"os"
	"github.com/say-bonjour/be-say-bonjour/security"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = "jgnfasdmfksd"

func CreateToken(principal security.Principal) (map[string]interface{}, error) {
	var err error
	os.Setenv("ACCESS_SECRET", SecretKey) //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["sub"] = 1
	atClaims["authorized"] = true
	atClaims["principal"] = principal
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	rt := jwt.New(jwt.SigningMethodHS256)
	rtClaims := rt.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["authorized"] = true
	rtClaims["principal"] = principal
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshToken, err := rt.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"accessToken":  token,
		"refreshToken": refreshToken,
		"principal":    principal,
	}, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
}
