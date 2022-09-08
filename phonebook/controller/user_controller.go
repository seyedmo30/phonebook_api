package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/seyedmo30/phonebook_api/phonebook/models"
	"github.com/seyedmo30/phonebook_api/phonebook/services"
)

var validate *validator.Validate

func MainPage(c *gin.Context) {
	//TODO bind
	//TODO validation

	user_service := services.User{

		Username: "aa",
	}
	user, _ := user_service.Exist_by_username()

	c.JSON(200, gin.H{
		"message": user,
	})
}

func Get_users(c *gin.Context) {

	user_service := services.User{

		Username: "aa",
	}
	user, _ := user_service.Exist_by_username()

	c.JSON(200, gin.H{
		"message": user,
	})
}

func Get_total_user(c *gin.Context) {

	arg := c.PostForm("state")
	println(arg)
	c.JSON(200, gin.H{
		"message": "true",
	})

}
func Get_user_by_username(c *gin.Context) {
	var username string

	username = c.Param("username")
	user, _ := models.Get_user_by_username(&username)

	c.JSON(200, gin.H{
		"message": user,
	})

}
func Test(c *gin.Context) {

	gg := c.Request.URL.Query()

	c.JSON(200, gin.H{
		"Accepted": gg,
		"Request":  c.Keys,
		"sss":      c.Params,
		"aaaa":     c.Writer,
	})
}
