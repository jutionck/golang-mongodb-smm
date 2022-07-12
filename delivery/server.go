package delivery

import (
	"enigmacamp.com/golang-with-mongodb/config"
	"enigmacamp.com/golang-with-mongodb/delivery/controller"
	"enigmacamp.com/golang-with-mongodb/manager"
	"fmt"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.useCaseManager.ProductRegistrationUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func NewServer() *appServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
	}
}
