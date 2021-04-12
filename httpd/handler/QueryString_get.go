package handler

import "github.com/gin-gonic/gin"

func QueryStrings(c *gin.Context) {
	name := c.Query("name") // Query = mengambil data dari data yang dilempar pada query
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
