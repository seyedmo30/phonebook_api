package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seyedmo30/phonebook_api/phonebook/services"
)

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(120)"`
	Number string `form:"number" valid:"Required;MaxSize(10)"`
	
}

func AddTag(c *gin.Context) {
	// var (
	// 	appG = app.Gin{C: c}
	// 	form AddTagForm
	// )

	// httpCode, errCode := app.BindAndValid(c, &form)
	// if errCode != e.SUCCESS {
	// 	appG.Response(httpCode, errCode, nil)
	// 	return
	// }

	contactService := contact_service.Contact{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}










	exists, err := tagService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
