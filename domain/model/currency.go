package model

import (
	"github.com/google/uuid"
	"time"
)

type Currency struct {
	ID         uuid.UUID `json:"id"`
	Code       string    `json:"code"`
	Value      float64   `json:"value"`
	CreateDate time.Time `json:"createDate"`
}

type Currencies []Currency

type Meta struct {
	LastUpdatedAt time.Time `json:"last_updated_at"`
}

type Data []CurrencyValue

type CurrencyValue struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type ResponseCurrency struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}
