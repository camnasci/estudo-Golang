package controller

import (
	"modulo/model"
	"modulo/service"
	"sync"

	"github.com/gin-gonic/gin"
)

func MyapiController(c *gin.Context) {

	type chanResp struct {
		model *model.Response
		err   error
	}

	currencyIDS := []string{"bitcoin", "ethereum", "monero"}
	currencyMap := make(map[string]*model.Response)
	wg := sync.WaitGroup{}

	channel := make(chan chanResp)
	wg.Add(len(currencyIDS))
	for _, id := range currencyIDS {
		go func(id string) {
			defer wg.Done()

			currency, err := service.GetCurrency(id)
			channel <- chanResp{
				model: currency,
				err:   err,
			}
		}(id)
	}

	go func() {
		defer close(channel)
		wg.Wait()
	}()

	status := 200
	for resp := range channel {
		if resp.err != nil {
			status = 206
		}
		currencyMap[resp.model.ID] = resp.model
	}

	c.JSON(status, currencyMap)
}
