package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seyedmo30/phonebook_api/phonebook/services"
)

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

func Test(c *gin.Context) {

	gg := c.Request.URL.Query()

	dd, ll := c.GetQuery("salam")
	println(dd, ll)
	println()
	println(gg)
	c.JSON(200, gin.H{
		"Accepted": gg,
		"Request":  c.Keys,
		"sss":      c.Params,
		"aaaa":     c.Writer,
	})
}
