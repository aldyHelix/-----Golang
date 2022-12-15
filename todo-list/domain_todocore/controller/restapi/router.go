package restapi

import (
	"todo-list/shared/gogen"
	"todo-list/shared/infrastructure/config"
	"todo-list/shared/infrastructure/logger"
	"todo-list/shared/infrastructure/token"

	"github.com/gin-gonic/gin"
)

type selectedRouter = gin.IRouter

type ginController struct {
	*gogen.BaseController
	log      logger.Logger
	cfg      *config.Config
	jwtToken token.JWTToken
}

func NewGinController(log logger.Logger, cfg *config.Config, tk token.JWTToken) gogen.RegisterRouterHandler[selectedRouter] {
	return &ginController{
		BaseController: gogen.NewBaseController(),
		log:            log,
		cfg:            cfg,
		jwtToken:       tk,
	}
}

func (r *ginController) RegisterRouter(router selectedRouter) {

	resource := router.Group("/api/v1", r.authentication())
	resource.GET("/todo", r.authorization(), r.getAllTodoHandler())
	resource.PUT("/todo/:todo_id", r.authorization(), r.runTodoCheckHandler())
	resource.POST("/todo", r.authorization(), r.runTodoCreateHandler())

}
