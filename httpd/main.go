package main

import (
	"belajar-gin/httpd/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	// default statement for gin
	r := gin.Default()

	// routes
	r.GET("/", handler.HomePage)
	r.POST("/", handler.PostHomePage) 				// mengambil data yang dimasukkan dari Body->raw postman
	r.GET("/query", handler.QueryStrings)             // /query?name=hoiril&age=24
	r.GET("/path/:name/:age", handler.PathParameters) // /path/hoiril/24

	// run statement
	r.Run()
}

// NOTE: tutorial by: https://www.youtube.com/watch?v=RkmvVFZJJvs
