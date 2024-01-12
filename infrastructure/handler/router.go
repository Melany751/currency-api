package handler

import (
	"challenge/domain/model"
	routerCurrency "challenge/infrastructure/handler/currency"
)

func InitRoutes(specification model.RouterSpecification) {
	routerCurrency.NewRouter(specification)
}
