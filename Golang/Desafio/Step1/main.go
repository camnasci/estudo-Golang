package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/myapi", MyapiController)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
	}
}
