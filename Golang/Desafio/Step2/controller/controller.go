package controller

import (
	"modulo/service"

	"github.com/gin-gonic/gin"
)

func MyapiController(c *gin.Context) {
	currencyID := c.Request.URL.Query().Get("ID")
	if currencyID == "" {
		currencyID = "bitcoin"
	}

	resp, err := service.GetCurrencyValue(currencyID)
	if err != nil {
		c.JSON(206, resp)
		return
	}

	c.JSON(200, resp)
}
