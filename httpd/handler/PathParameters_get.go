package handler

import "github.com/gin-gonic/gin"

func PathParameters(c *gin.Context) {
	name := c.Param("name") // Param = mengambil data dari path
	age := c.Param("age")

	c.JSON(200, gin.H {
		"name" : name,
		"age" : age,
	})
}
