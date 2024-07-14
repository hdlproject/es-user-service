package interactor

import (
	"errors"

	"github.com/hdlproject/es-user-service/entity"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/input_port"
	"github.com/hdlproject/es-user-service/use_case/output_port"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	TopUp struct {
		userRepo       output_port.UserRepo
		topUpEventRepo output_port.TopUpEventRepo
	}
)

func NewTopUpUseCase(userRepo output_port.UserRepo,
	topUpEventRepo output_port.TopUpEventRepo) *TopUp {

	return &TopUp{
		userRepo:       userRepo,
		topUpEventRepo: topUpEventRepo,
	}
}

func (instance *TopUp) TopUp(request input_port.TopUpRequest) (response input_port.TopUpResponse, err error) {
	topUpEvent := entity.TopUpEvent{
		TransactionEventID: request.TransactionEventID,
		UserID:             request.UserID,
		Amount:             request.Increment,
	}
	isAlreadyApplied, err := instance.topUpEventRepo.IsAlreadyApplied(topUpEvent)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return input_port.TopUpResponse{}, helper.WrapError(err)
	}
	if isAlreadyApplied {
		return input_port.TopUpResponse{}, &helper.RequestAlreadyProcessed{
			Message: "top up request already processed",
		}
	}

	err = instance.userRepo.IncreaseBalance(request.UserID, request.Increment)
	if err != nil {
		return input_port.TopUpResponse{}, helper.WrapError(err)
	}

	_, err = instance.topUpEventRepo.Insert(topUpEvent)
	if err != nil {
		return input_port.TopUpResponse{}, helper.WrapError(err)
	}

	return input_port.TopUpResponse{
		Ok:      true,
		Message: success,
	}, nil
}
