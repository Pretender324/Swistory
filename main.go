package main

import (
	"swistory/controller"
	"swistory/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)

	menuEngine := engine.Group("/menu")
	{
		v1 := menuEngine.Group("/v1")
		{
			v1.POST("/add", controller.MenuAdd)
			v1.GET("/list", controller.MenuList)
		}
	}
	engine.Run(":5000") // listen and serve on 0.0.0.0:8080
}
