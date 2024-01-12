package api

import (
	"challenge/domain/model"
	"time"
)

type CurrencyApi interface {
	GetCurrenciesApi(apiKey string, timeOut int) (*time.Time, *model.Data, error)
}
