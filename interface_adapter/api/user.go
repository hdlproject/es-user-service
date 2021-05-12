package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hdlproject/es-user-service/config"
	"github.com/hdlproject/es-user-service/interface_adapter/database"
	"github.com/hdlproject/es-user-service/use_case/interactor"
	"log"
	"net/http"
)

type (
	UserController struct {
		userService *userService
	}
)

func RegisterUserAPI() {
	configInstance, _ := config.GetInstance()

	httpServer := GetHTTPServer(configInstance.Port)

	postgresClient, _ := database.GetPostgresClient(configInstance.Database)

	userRouter := httpServer.apiRouter.Group("/user")
	userRouter.POST("/register", NewUserController(
		interactor.NewRegisterUseCase(
			database.NewUserRepo(postgresClient),
		),
	).Register)
}

func NewUserController(registerUseCase *interactor.Register) *UserController {
	return &UserController{
		userService: NewUserService(
			registerUseCase,
		),
	}
}

func (instance *UserController) Register(ctx *gin.Context) {
	request, err := registerRequest{}.parse(ctx)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, registerResponse{Ok: false, Message: parseRequestFailure})
		return
	}

	response, err := instance.userService.register(request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, registerResponse{Ok: false, Message: defaultProcessError})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
