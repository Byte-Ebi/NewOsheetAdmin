package authentication

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type authClaims struct {
	jwt.StandardClaims
	Account string `json:"account"`
}

var jwtKey []byte

func init() {
	viper.SetConfigFile("configs/app.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig()failed,err:%v\n", err)
		return
	}

	jwtKey = []byte(viper.GetString("authentication.jwtKey"))
}

func verifyToken(tokenString string) (string, int64, error) {
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return "", 0, err
	}
	if !token.Valid {
		return "", 0, fmt.Errorf("invalid token")
	}

	return claims.Account, claims.ExpiresAt, nil
}

// TODO 註銷 token

// TODO 產生 token
