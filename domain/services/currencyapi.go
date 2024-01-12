package services

type UseCaseCurrencyApi interface {
	GetCurrenciesValues(apiKey string, interval int, timeOut int) error
}
