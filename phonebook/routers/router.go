package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/seyedmo30/phonebook_api/phonebook/controller"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/")
	{

		// api.GET("/users", controller.Get_users)
		api.GET("/", controller.MainPage)
		api.GET("/test", controller.Test)
		// api.GET("/users/total", controller.Get_total_user)
		api.POST("/contact/add", controller.AddContact)
		api.GET("/users/:username", controller.Get_user_by_username)
		api.GET("/contact/search", controller.SearchContact)
		api.GET("/contact/:id/", controller.GetContact)
		api.GET("/contact", controller.ListContact)

	}

	return r
}
