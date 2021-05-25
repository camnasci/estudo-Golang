package main

import (
	"fmt"
	"modulo/client"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/myapi", client.MyapiController)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
	}
}
