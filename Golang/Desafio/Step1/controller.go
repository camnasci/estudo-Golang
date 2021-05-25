package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Text string `json:"data"`
}

func MyapiController(c *gin.Context) {
	param := c.Request.URL.Query().Get("data")
	if param == "" {
		c.JSON(500, "query param requird")
		return
	}
	resp := Response{
		Text: param,
	}
	c.JSON(200, resp)
}
