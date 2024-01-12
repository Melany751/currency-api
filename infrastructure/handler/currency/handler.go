package currency

import (
	"challenge/domain/model"
	"challenge/domain/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

type handler struct {
	useCase services.UseCaseCurrency
}

func newHandler(useCase services.UseCaseCurrency) handler {
	return handler{useCase}
}

func (h handler) get(c *gin.Context) {

	currencyParam := c.Param("CURRENCY")

	regex := regexp.MustCompile(`^/currencies/([a-zA-Z][a-zA-Z])([a-zA-Z]|[a-zA-Z]\?finit=([a-zA-Z0-9:\-]+)|[a-zA-Z]\?fend=([a-zA-Z0-9:\-]+)|[a-zA-Z]\?finit=([a-zA-Z0-9:\-]+)&fend=([a-zA-Z0-9:\-]+))$`)

	if !regex.MatchString(c.Request.URL.String()) {
		fmt.Println("La URL no es válida")
		c.JSON(services.BadRequest(model.ResponseError{"La URL no es valida"}))
		return
	}

	finit := c.Query("finit")

	if finit == "" {
		finit = "2000-01-01 00:00:00"
	}

	fend := c.Query("fend")

	if fend == "" {
		fend = "9999-01-01 00:00:00"
	}

	m, err := h.useCase.Get(currencyParam, finit, fend)
	if err != nil {
		c.JSON(services.SystemError(model.ResponseError{Error: err.Error()}))
		return
	}
	c.JSON(services.OK(m))
}
