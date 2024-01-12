package currency

import (
	useCaseCurrency "challenge/application/usecase/currency"
	"challenge/domain/model"
	storageCurrency "challenge/infrastructure/storage/postgres"
	"github.com/gin-gonic/gin"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := useCaseCurrency.New(storageCurrency.NewCurrencyStorage(specification.DB))

	return newHandler(useCase)
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("currencies", middlewares...)

	routes.GET("/:CURRENCY", h.get)
}
