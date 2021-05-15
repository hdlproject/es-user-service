package interactor

import (
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/output_port"
)

type (
	TopUpRequest struct {
		UserID    uint
		Increment uint64
	}

	TopUpResponse struct {
		Ok      bool
		Message string
	}

	TopUp struct {
		userRepo output_port.UserRepo
	}
)

func NewTopUpUseCase(userRepo output_port.UserRepo) *TopUp {
	return &TopUp{
		userRepo: userRepo,
	}
}

func (instance *TopUp) TopUp(request TopUpRequest) (response TopUpResponse, err error) {
	err = instance.userRepo.IncreaseBalance(request.UserID, request.Increment)
	if err != nil {
		return TopUpResponse{}, helper.WrapError(err)
	}

	return TopUpResponse{
		Ok:      true,
		Message: success,
	}, nil
}
