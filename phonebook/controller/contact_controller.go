package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/seyedmo30/phonebook_api/phonebook/entity"
	"github.com/seyedmo30/phonebook_api/phonebook/services"
)

type AddContactForm struct {
	Name   string `form:"name"  json:"name"  binding:"required" `
	Number int64  `form:"number"  json:"number"  binding:"required"`
}

type SearchContactForm struct {
	Number int    `json:"number"  binding:"required" `
	Name   string `json:"name"  binding:"required" validate:"min=2,max=40"`
}

func SearchContact(c *gin.Context) {
	//TODO bind
	var s SearchContactForm
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

	contact, err := services.SearchContact(&s.Name)

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

func GetContact(c *gin.Context) {

	id_string := c.Param("id")
	id, err_atio := strconv.Atoi(id_string)
	if err_atio != nil {
		fmt.Println(err_atio)
		c.JSON(500, gin.H{
			"message": err_atio,
		})
		return
	}

	validate = validator.New()
	err := validate.Var(id, "required,gte=1")

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

	contact, err := services.GetContact(&id)

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

func AddContact(c *gin.Context) {
	var json AddContactForm

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var contact entity.Contact
	contact.Name = json.Name
	contact.Number = json.Number

	err := services.AddContact(&contact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})

}

func ListContact(c *gin.Context) {
	//TODO bind
	page_string, _ := c.GetQuery("page")
	page, _ := strconv.Atoi(page_string)

	println("page", page)
	validate = validator.New()
	err := validate.Var(page, "gte=0")

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

	if page == 0 {
		page = 10
	}

	contact, err := services.ListContact(&page)

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
