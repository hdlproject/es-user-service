package api

import (
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/interactor"
)

type (
	userService struct {
		registerUseCase *interactor.Register
	}
)

func NewUserService(registerUseCase *interactor.Register) *userService {
	return &userService{
		registerUseCase: registerUseCase,
	}
}

func (instance *userService) register(request registerRequest) (registerResponse, error) {
	useCaseRequest := request.getUseCase()

	useCaseResponse, err := instance.registerUseCase.Register(useCaseRequest)
	if err != nil {
		return registerResponse{}, helper.WrapError(err)
	}

	return registerResponse{}.fromUseCase(useCaseResponse), nil
}
