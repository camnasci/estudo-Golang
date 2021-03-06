package main

import (
	"fmt"
	"modulo/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/myapi", controller.MyapiController)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
	}
}
