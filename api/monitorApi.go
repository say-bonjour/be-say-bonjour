package api

import (
	"github.com/say-bonjour/be-say-bonjour/model"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	db := getDatabase(c)
	users := make([]model.User, 0)

	if err := db.Find(&users).Limit(10).Error; err != nil {
		internalServerError(c, err.Error())
		return
	}

	ok(c, "OK")
}

func Panic(c *gin.Context) {
	panic("Test")
}
