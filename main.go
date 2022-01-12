package main

import (
	"app/handler"
	"app/lib"
	"app/model"

	"github.com/gin-gonic/gin"
)

func main() {
	lib.DBOpen()
	defer lib.DBClose()
	lib.GetDBConn().DB.AutoMigrate(&model.Person{})

	r := gin.Default()
	nyt := r.Group("person")
	{
		nyt.GET(":id", handler.GetPerson())
		nyt.POST("", handler.PostPerson())
		nyt.PUT("", handler.PutPerson())
		nyt.DELETE("", handler.DeletePerson())
	}

	r.Run("0.0.0.0:8080")

}
