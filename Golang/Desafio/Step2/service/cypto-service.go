package service

import (
	"encoding/json"
	"errors"
	"modulo/model"
	"net/http"
)

func GetCurrencyValue(currencyID string) (*model.Response, error) {
	respFinal := model.Response{
		ID:      currencyID,
		Partial: true,
	}

	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + currencyID + "&vs_currencies=usd")
	if err != nil {
		return &respFinal, err
	}

	defer resp.Body.Close()

	// Otra forma
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(resp.Body)

	// var cResp model.Cryptoresponse
	// err = json.Unmarshal(buf.Bytes(), cResp)
	// if err != nil {
	// 	log.Fatal("ooopsss! an error occurred, please try again")
	// }

	var cResp model.CurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return &respFinal, err
	}

	value, ok := cResp[currencyID]
	if !ok {
		return &respFinal, errors.New("Currency not found")
	}

	respFinal.Partial = false
	respFinal.Content = &model.Content{
		Price:    value.Usd,
		Currency: "USD",
	}

	return &respFinal, nil
}
