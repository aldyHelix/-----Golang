package application

import (
	"todo-list/domain_todocore/controller/restapi"
	"todo-list/domain_todocore/gateway/withgorm"
	"todo-list/domain_todocore/usecase/getalltodo"
	"todo-list/domain_todocore/usecase/runtodocheck"
	"todo-list/domain_todocore/usecase/runtodocreate"
	"todo-list/shared/gogen"
	"todo-list/shared/infrastructure/config"
	"todo-list/shared/infrastructure/logger"
	"todo-list/shared/infrastructure/server"
	"todo-list/shared/infrastructure/token"
)

type todoapp struct{}

func NewTodoapp() gogen.Runner {
	return &todoapp{}
}

func (todoapp) Run() error {

	const appName = "todoapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := restapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		getalltodo.NewUsecase(datasource),
		runtodocheck.NewUsecase(datasource),
		runtodocreate.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
