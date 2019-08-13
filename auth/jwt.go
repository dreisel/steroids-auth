package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	users2 "github.com/dreisel/steroids-auth/users"
	"github.com/dreisel/steroids-auth/utils"
	"strconv"
	"time"
)

var jwtKey = []byte(utils.GetEnv("JWT_SECRET_KEY", "jwt-secret"))
var jwtIssuer = utils.GetEnv("JWT_ISSUER", "todos-on-steroids")

type Token struct {
	Value   string
	Expires time.Time
}

func CreateToken(user users2.User) (Token, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		Issuer:    jwtIssuer,
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return Token{
		Value:   tokenString,
		Expires: expirationTime,
	}, err
}

func getSecret(_ *jwt.Token) (interface{}, error) {
	return jwtKey, nil
}

func ParseToken(token string) (jwt.StandardClaims, error) {
	claims := jwt.StandardClaims{}
	fmt.Println("token ", token)
	tkn, err := jwt.ParseWithClaims(token, &claims, getSecret)
	if err != nil {
		return claims, err
	}
	if !tkn.Valid || claims.Issuer != jwtIssuer {
		return claims, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
