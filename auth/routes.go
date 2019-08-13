package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dreisel/steroids-auth/users"
	"github.com/dreisel/steroids-auth/utils"
	"github.com/gin-gonic/gin"
	"time"
)

var authCookieName = utils.GetEnv("AUTH_COOKIE_NAME", "auth")
var authCookieDomain = utils.GetEnv("AUTH_COOKIE_DOMAIN", "")
var authCookieSecure = utils.GetEnv("AUTH_COOKIE_SECURE", "false") == "true"

type authCookie struct {
}

func (ac authCookie) set(c *gin.Context, user users.User) error {
	token, err := CreateToken(user)
	if err != nil {
		return err
	}
	ac.setCookie(c, token)
	return nil
}

func (ac authCookie) setCookie(c *gin.Context, token Token) {
	c.SetCookie(authCookieName, token.Value, int(token.Expires.Sub(time.Now()).Seconds()), "/", authCookieDomain, authCookieSecure, true)
}

func (ac authCookie) delete(c *gin.Context) {
	ac.setCookie(c, Token{
		Value:   "",
		Expires: time.Now(),
	})
}

func (ac authCookie) get(c *gin.Context) (jwt.StandardClaims, error) {
	token, err := c.Cookie(authCookieName)
	if err != nil {
		return jwt.StandardClaims{}, err
	}
	fmt.Println(token)
	return ParseToken(token)
}
