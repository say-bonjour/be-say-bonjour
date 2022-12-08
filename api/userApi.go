package api

import (
	"github.com/say-bonjour/be-say-bonjour/model"
	"github.com/say-bonjour/be-say-bonjour/utility"

	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	db := getDatabase(c)
	users := make([]model.User, 0)
	db.Find(&users)

	for i := range users {
		users[i].Password = "-"
	}

	ok(c, users)
}

func Register(c *gin.Context) {
	var request model.User
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err.Error())
		return
	}

	user := model.User{Name: request.Name, Email: request.Email}

	password, err := utility.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Illegal Password")
		return
	}
	user.Password = password

	db := getDatabase(c)
	if err := db.Create(&user).Error; err != nil {
		internalServerError(c, err.Error())
		return
	}

	ok(c, user)
}

func FindUser(c *gin.Context) {
	var user model.User

	db := getDatabase(c)
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		badRequest(c, "Record not found!")
		return
	}

	ok(c, user)
}

func MyUser(c *gin.Context) {
	var user model.User

	principal := getPrincipal(c)

	db := getDatabase(c)
	if err := db.Where("id = ?", principal.Id).First(&user).Error; err != nil {
		badRequest(c, "Record not found!")
		return
	}

	ok(c, user)
}

func UpdateUser(c *gin.Context) {
	var request model.User
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err.Error())
		return
	}

	db := getDatabase(c)
	var user model.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		badRequest(c, "Record not found!")
		return
	}

	user.Name = request.Name
	user.Email = request.Email

	db.Save(&user)
	ok(c, user)
}

func DeleteUser(c *gin.Context) {
	db := getDatabase(c)
	var book model.User
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		badRequest(c, "Record not found!")
		return
	}

	db.Delete(&book)

	ok(c, gin.H{"data": true})
}
