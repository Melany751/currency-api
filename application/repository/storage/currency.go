package storage

import (
	"challenge/domain/model"
)

type CurrencyStorage interface {
	GetCurrenciesStorage(currencyParam string, finit string, fend string) (model.Currencies, error)
	CreateCurrencyStorage(currency model.Currency) error
}
