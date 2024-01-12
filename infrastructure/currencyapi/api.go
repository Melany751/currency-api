package currencyapi

import (
	"challenge/domain/model"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type CurrencyApi struct{}

func NewApi(currencyApi CurrencyApi) CurrencyApi {
	return currencyApi
}

func (c CurrencyApi) GetCurrenciesApi(apiKey string, timeOut int) (*time.Time, *model.Data, error) {
	var currencyModel model.ResponseCurrency

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := "http://api.currencyapi.com/v3/latest?apikey=" + apiKey

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("Error al crear la solicitud:", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Verificar si el error fue causado por timeout
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return nil, nil, fmt.Errorf("La solicitud excedió el tiempo de espera")
		}
		return nil, nil, fmt.Errorf("Error al realizar la solicitud:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("Error al leer la respuesta:", err)
	}

	err = json.Unmarshal(body, &currencyModel)
	if err != nil {
		return nil, nil, fmt.Errorf("Error al convertir el modelo:", err)
	}

	return &currencyModel.Meta.LastUpdatedAt, &currencyModel.Data, nil
}
