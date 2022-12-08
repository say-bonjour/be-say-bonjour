package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/say-bonjour/be-say-bonjour/constant"
	"github.com/say-bonjour/be-say-bonjour/core"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.Path, "/api/public") {
			authHeader := c.GetHeader("Authorization")
			token, err := core.ValidateToken(authHeader)
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				c.Set(constant.Gin_Principal, claims["principal"])
			} else {
				log.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
