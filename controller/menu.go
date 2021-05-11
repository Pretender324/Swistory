package controller

import (
	"net/http"
	"swistory/model"
	"swistory/service"

	"github.com/gin-gonic/gin"
)

func MenuAdd(c *gin.Context) {
	menu := model.Menu{}
	err := c.Bind(&menu)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	menuService := service.MenuService{}
	err = menuService.SetMenu(&menu)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func MenuList(c *gin.Context) {
	menuService := service.MenuService{}
	MenuLists := menuService.GetMenuList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    MenuLists,
	})
}
