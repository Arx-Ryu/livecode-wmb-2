package delivery

import (
	"livecode-wmb-rest-api/config"
	"livecode-wmb-rest-api/delivery/controller"
	"livecode-wmb-rest-api/manager"
	"livecode-wmb-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager		manager.UseCaseManager
	engine 				*gin.Engine
	host				string
	token 				utils.Token
}

func Server() *appServer {
	r := gin.Default()
	c := config.NewConfig()
	infra := manager.NewInfra(c)	
	repoManager := manager.NewRepoManager(infra)
	useCaseManager := manager.NewUseCaseManager(repoManager.BillRepo(), repoManager.CustomerRepo(), repoManager.MenuRepo(), repoManager.LoginRepo())
	host := c.ApiConfig.Url
	return &appServer{
		useCaseManager: useCaseManager,
		engine: r,
		host: host,
		token: infra.Token(),
	}
}

func (a *appServer) initControllers() {
	controller.NewBillController(a.engine, a.useCaseManager, a.token)
	controller.NewMenuController(a.engine, a.useCaseManager)
	controller.NewCustController(a.engine, a.useCaseManager)
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

/*
set API_URL=localhost:8080
set API_HOST=localhost
set API_PORT=8080
*/