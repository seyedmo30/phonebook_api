package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/seyedmo30/phonebook_api/phonebook/services"
)

type AddContactForm struct {
	Name   string `form:"name"  json:"name"  binding:"required" `
	Number string `form:"number"  json:"number"  binding:"required"`
}

type SearchContactStruct struct {
	Number int    `json:"number"  binding:"required" validate:"gte=1000"`
	Name   string `json:"name"  binding:"required" validate:"min=3,max=4"`
}

func SearchContact(c *gin.Context) {
	//TODO bind
	var s SearchContactStruct
	s.Name, _ = c.GetQuery("name")
	s.Number = c.GetInt("number")
	println(s.Number)
	validate = validator.New()
	err := validate.Struct(s)
	// errs := validate.Var(name, "required,min=3,max=45")
	// errs = validate.Var(name, "required,min=3,max=45")

	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	name := "hasan"
	contact, err := services.SearchContact(&name)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"results": contact,
	})

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
