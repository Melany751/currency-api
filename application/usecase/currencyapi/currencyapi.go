package currencyapi

import (
	"challenge/application/repository/api"
	"challenge/application/repository/storage"
	"challenge/domain/model"
	"github.com/google/uuid"
	"time"
)

type CurrencyApi struct {
	api             api.CurrencyApi
	currencyStorage storage.CurrencyStorage
	registryStorage storage.RegistryStorage
}

func New(api api.CurrencyApi, currencyStorage storage.CurrencyStorage, registryStorage storage.RegistryStorage) CurrencyApi {
	return CurrencyApi{api, currencyStorage, registryStorage}
}

func (u CurrencyApi) GetCurrenciesValues(apiKey string, interval int, timeOut int) error {

	for true {
		//var registry model.Registry

		//startDate := time.Now()

		dateCreate, currencies, err := u.api.GetCurrenciesApi(apiKey, timeOut)
		if err != nil {
			return err
		}

		err = u.saveCurrenciesValues(*dateCreate, *currencies)
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(interval) * time.Minute)
	}

	return nil
}

func (u CurrencyApi) saveCurrenciesValues(createDate time.Time, currencies model.Data) error {
	var currencyResult = model.Currency{}
	for _, currency := range currencies {
		newId, err := uuid.NewUUID()
		if err != nil {
			return err
		}

		currencyResult.ID = newId
		currencyResult.Code = currency.Code
		currencyResult.Value = currency.Value
		currencyResult.CreateDate = createDate

		err = u.currencyStorage.CreateCurrencyStorage(currencyResult)
		if err != nil {
			return err
		}
	}

	return nil
}
