package main

import (
	"github.com/say-bonjour/be-say-bonjour/constant"
	"github.com/say-bonjour/be-say-bonjour/core"
	"github.com/say-bonjour/be-say-bonjour/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	properties, err := core.LoadConfiguration("application.json")
	if err != nil {
		panic("Could not parse properties file: " + err.Error())
	}

	db := core.SetupDatabase(properties.Database)

	migrateDB(db)

	app := gin.Default()

	app.Use(func(c *gin.Context) {
		c.Set(constant.Gin_Database, db)
		c.Set(constant.Gin_Properties, &properties)
	})
	app.Use(middleware.AuthorizeJWT())

	setupRoutes(app)
	app.Run(":" + properties.Port)
}
