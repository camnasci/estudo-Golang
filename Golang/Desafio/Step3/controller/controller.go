package controller

import (
	"modulo/model"
	"modulo/service"
	"sync"

	"github.com/gin-gonic/gin"
)

func MyapiController(c *gin.Context) {

	currencyIDS := []string{"bitcoin", "ethereum", "monero"}

	currencyMap := make(map[string]*model.Response)
	wg := sync.WaitGroup{}

	for _, id := range currencyIDS {
		wg.Add(1)
		go func(id string) {
			currency, err := service.GetCurrency(id)

			if err != nil {
				return
			}

			currencyMap[id] = currency
			wg.Done()
		}(id)
	}
	wg.Wait()

	c.JSON(200, currencyMap)

}
