package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func PostHomePage(c *gin.Context) {

	body := c.Request.Body				// ambil data dari Request Body
	value, err := ioutil.ReadAll(body)  // ioutil package untuk I/O yang ditampung di value, menggunakan data body
	if err != nil {						// cek error
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H {
		"message" : string(value),  	// menampilkan apa yang diinputkan pada body
	})
}
