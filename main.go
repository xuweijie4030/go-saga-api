package main

import (
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/dbdao/gormdb"
	"github.com/carefreex-io/logger"
	"github.com/carefreex-io/xhttp"
	"gosagaapi/app/daemon"
	"gosagaapi/app/middleware"
	"gosagaapi/app/validation"
	"gosagaapi/router"
)

func main() {
	config.InitConfig()

	logger.InitLogger()

	if err := gormdb.InitDB(); err != nil {
		logger.Fatalf("mysql.InitDB failed: err=%v", err)
	}

	daemon.RunStartBeforeFn()

	http := xhttp.NewXHttp()

	validation.RegisterValidation()

	middleware.Register(http.Engine)

	router.Register(http.Engine)

	http.Start()

	daemon.RunStoppedFn()
}
