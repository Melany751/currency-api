package bootstrap

import (
	"challenge/domain/model"
	"challenge/infrastructure/handler"
	"context"
	"github.com/joho/godotenv"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
)

func Run(boot []byte) {
	_ = godotenv.Load()

	db := newDatabase()

	ginEntry := newGinEntry(boot)
	ginEntry.Bootstrap(context.Background())

	api := ginEntry.Router

	handler.InitRoutes(model.RouterSpecification{
		Api: api,
		DB:  db,
	})

	rkentry.GlobalAppCtx.WaitForShutdownSig()
	ginEntry.Interrupt(context.Background())

	/*interval_min := os.Getenv("INTERVAL_MIN")
	intInterval, _ := strconv.Atoi(interval_min)

	timeOutIn := os.Getenv("TIMEOUT")
	intTimeOut, _ := strconv.Atoi(timeOutIn)*/

	//c.useCase.GetCurrenciesValues(os.Getenv("API_KEY"), intInterval, intTimeOut)
}
