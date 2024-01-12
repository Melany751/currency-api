package currency

import (
	"challenge/application/repository/storage"
	"challenge/domain/model"
	"fmt"
)

type Currency struct {
	storage storage.CurrencyStorage
}

func New(storage storage.CurrencyStorage) Currency {
	return Currency{storage}
}

func (u Currency) Get(currencyParam string, fInit string, fEnd string) (model.Currencies, error) {

	currencies, err := u.storage.GetCurrenciesStorage(currencyParam, fInit, fEnd)
	if err != nil {
		return nil, fmt.Errorf("currency.storage.Get(): %w", err)
	}

	return currencies, nil
}
