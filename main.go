package main

import (
	"swistory/controller"
	"swistory/db"
	"swistory/middleware"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)
	db.Init()

	menuEngine := engine.Group("/menu")
	{
		v1 := menuEngine.Group("/v1")
		{
			v1.POST("/add", controller.MenuAdd)
			v1.GET("/list", controller.MenuList)
			v1.PUT("/update", controller.MenuUpdate)
			v1.DELETE("/delete", controller.MenuDelete)
		}
	}
	engine.Run(":5000") // listen and serve on 0.0.0.0:8080
}
