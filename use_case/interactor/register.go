package interactor

import (
	"github.com/hdlproject/es-user-service/entity"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/input_port"
	"github.com/hdlproject/es-user-service/use_case/output_port"
)

type (
	Register struct {
		userRepo output_port.UserRepo
	}
)

func NewRegisterUseCase(userRepo output_port.UserRepo) *Register {
	return &Register{
		userRepo: userRepo,
	}
}

func (instance *Register) Register(request input_port.RegisterRequest) (response input_port.RegisterResponse, err error) {
	user := entity.User{
		Auth: entity.UserAuth{
			Username: request.Username,
			Password: request.Password,
		},
	}
	userID, err := instance.userRepo.Register(user)
	if err != nil {
		return input_port.RegisterResponse{}, helper.WrapError(err)
	}

	return input_port.RegisterResponse{
		UserID: userID,
	}, nil
}
