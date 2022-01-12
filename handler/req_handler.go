package handler

import (
	"app/lib"
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPerson() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var person model.Person

		db := lib.GetDBConn().DB
		db.First(&person, id)
		c.JSON(http.StatusOK, person)
	}
}

func PostPerson() gin.HandlerFunc {
	return func(c *gin.Context) {
		var person model.Person
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := lib.GetDBConn().DB
		db.Create(&person)
		c.JSON(http.StatusOK, person)
	}
}

func PutPerson() gin.HandlerFunc {
	return func(c *gin.Context) {
		var person model.Person
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := lib.GetDBConn().DB
		db.Save(&person)
		c.JSON(http.StatusOK, person)
	}
}

func DeletePerson() gin.HandlerFunc {
	return func(c *gin.Context) {
		var person model.Person
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := lib.GetDBConn().DB
		db.Delete(&person)
		c.JSON(http.StatusOK, person)
	}
}
