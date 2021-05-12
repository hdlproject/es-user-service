package interactor

import (
	"github.com/hdlproject/es-user-service/entity"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/output_port"
)

type (
	RegisterRequest struct{}

	RegisterResponse struct {
		Ok      bool
		Message string
		UserID  uint
	}

	Register struct {
		userRepo output_port.UserRepo
	}
)

func NewRegisterUseCase(userRepo output_port.UserRepo) *Register {
	return &Register{
		userRepo: userRepo,
	}
}

func (instance *Register) Register(request RegisterRequest) (response RegisterResponse, err error) {
	user := entity.User{}
	userID, err := instance.userRepo.Register(user)
	if err != nil {
		return RegisterResponse{}, helper.WrapError(err)
	}

	return RegisterResponse{
		UserID: userID,
	}, nil
}
