package api

import (
	"encoding/json"
	"net/http"

	//"github.com/say-bonjour/be-say-bonjour/constant"
	"github.com/say-bonjour/be-say-bonjour/constant"
	"github.com/say-bonjour/be-say-bonjour/security"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func getProperties(c *gin.Context) *core.Properties {
// 	return c.MustGet(constant.Gin_Properties).(*core.Properties)
// }

func getDatabase(c *gin.Context) *gorm.DB {
	return c.MustGet(constant.Gin_Database).(*gorm.DB)
}

func getPrincipal(c *gin.Context) security.Principal {
	principalMap := c.MustGet(constant.Gin_Principal)
	return convertToPrincipal(principalMap)
}

func convertToPrincipal(dataMap interface{}) security.Principal {
	var principal security.Principal

	jsonStr, err := json.Marshal(dataMap)
	if err != nil {
		panic("Failed to get principal: " + err.Error())
	}

	if err := json.Unmarshal(jsonStr, &principal); err != nil {
		panic("Failed to get principal: " + err.Error())
	}

	return principal
}

func badRequest(c *gin.Context, responseBody interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"error": responseBody})
}

func internalServerError(c *gin.Context, responseBody interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": responseBody})
}

func ok(c *gin.Context, responseBody interface{}) {
	c.JSON(http.StatusOK, responseBody)
}

// func created(c *gin.Context, responseBody interface{}) {
// 	c.JSON(http.StatusCreated, responseBody)
// }
