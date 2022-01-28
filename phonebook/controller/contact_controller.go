package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seyedmo30/phonebook_api/phonebook/services"
)

type AddContactForm struct {
	Name   string `form:"name"  json:"name"  binding:"required"`
	Number string `form:"number"  json:"number"  binding:"required"`
}

func Add_contact(c *gin.Context) {
	var json AddContactForm

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("con \n", json.Name)
	fmt.Println("con \n", json.Number)

	err := services.Add(&json.Name, &json.Number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})

}

// func AddTag(c *gin.Context) {
// 	// var (
// 	// 	appG = app.Gin{C: c}
// 	// 	form AddTagForm
// 	// )

// 	// httpCode, errCode := app.BindAndValid(c, &form)
// 	// if errCode != e.SUCCESS {
// 	// 	appG.Response(httpCode, errCode, nil)
// 	// 	return
// 	// }

// 	contactService := contact_service.Contact{
// 		Name:      form.Name,
// 		CreatedBy: form.CreatedBy,
// 		State:     form.State,
// 	}

// 	exists, err := tagService.ExistByName()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
// 		return
// 	}
// 	if exists {
// 		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
// 		return
// 	}

// 	err = tagService.Add()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
// 		return
// 	}

// 	appG.Response(http.StatusOK, e.SUCCESS, nil)
// }
