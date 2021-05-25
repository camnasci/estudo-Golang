package client

import (
	"encoding/json"
	"log"
	"modulo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MyapiController(c *gin.Context) {

	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}

	defer resp.Body.Close()

	var cResp model.Cryptoresponse

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	respFinal := model.Response{
		ID: "BTC",
		Content: model.Content{
			Price:    cResp.Bitcoin.Usd,
			Currency: "USD",
		},
		Partial: false,
	}

	c.JSON(200, respFinal)
}
