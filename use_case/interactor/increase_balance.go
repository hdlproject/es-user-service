package interactor

import (
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/output_port"
)

type (
	IncreaseBalanceRequest struct {
		UserID    uint
		Increment uint64
	}

	IncreaseBalanceResponse struct{}

	IncreaseBalance struct {
		userRepo output_port.UserRepo
	}
)

func (instance *IncreaseBalance) IncreaseBalance(request IncreaseBalanceRequest) (response IncreaseBalanceResponse, err error) {
	err = instance.userRepo.IncreaseBalance(request.UserID, request.Increment)
	if err != nil {
		return IncreaseBalanceResponse{}, helper.WrapError(err)
	}

	return IncreaseBalanceResponse{}, nil
}
