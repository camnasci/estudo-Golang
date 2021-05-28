package service

import (
	"encoding/json"
	"modulo/model"
	"net/http"
)

func GetCurrency(currencyID string) (*model.Response, error) {
	respFinal := model.Response{
		ID:      currencyID,
		Partial: true,
	}

	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + currencyID + "&vs_currencies=usd")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var cResp model.CurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return &respFinal, err
	}

	value, ok := cResp[currencyID]
	if !ok {
		return &respFinal, err
	}

	respFinal.Partial = false
	respFinal.Content = &model.Content{
		Price:    value.Usd,
		Currency: "USD",
	}

	return &respFinal, nil
}
