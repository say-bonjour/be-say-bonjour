package api

import (
	"github.com/say-bonjour/be-say-bonjour/core"
	"github.com/say-bonjour/be-say-bonjour/model"
	"github.com/say-bonjour/be-say-bonjour/security"
	"github.com/say-bonjour/be-say-bonjour/utility"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenReqBody struct {
	RefreshToken string `json:"refreshToken"`
}

func Login(c *gin.Context) {
	var request UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err.Error())
		return
	}

	var user model.User
	db := getDatabase(c)
	if err := db.Where("email = ?", request.Username).First(&user).Error; err != nil {
		badRequest(c, "Username Not Found")
		return
	}

	if !utility.CheckPasswordHash(request.Password, user.Password) {
		badRequest(c, "Username and password did not match")
		return
	}

	tokenMap, err := core.CreateToken(security.Principal{Id: user.Id, Name: user.Name, Email: user.Email, Role: model.Admin})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ok(c, tokenMap)
}

func RefreshToken(c *gin.Context) {
	var tokenReq TokenReqBody
	if err := c.ShouldBindJSON(&tokenReq); err != nil {
		badRequest(c, err.Error())
		return
	}

	token, err := core.ValidateToken(tokenReq.RefreshToken)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		tokenMap, err := core.CreateToken(convertToPrincipal(claims["principal"]))
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		ok(c, tokenMap)
	} else {
		badRequest(c, err.Error())
		return
	}
}
