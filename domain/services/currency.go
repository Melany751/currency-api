package services

import (
	"challenge/domain/model"
)

type UseCaseCurrency interface {
	Get(currencyParam string, fInit string, fEnd string) (model.Currencies, error)
}
