package messaging

import (
	"encoding/json"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/interactor"
)

type (
	topUpRequest struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Params struct {
			UserID uint   `json:"user_id"`
			Amount uint64 `json:"amount"`
		} `json:"params"`
	}

	topUpResponse struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
	}

	topUpService struct {
		topUpUseCase *interactor.TopUp
	}
)

func newTopUpService(topUpUseCase *interactor.TopUp) *topUpService {
	return &topUpService{
		topUpUseCase: topUpUseCase,
	}
}

func (instance *topUpService) topup(request topUpRequest) (topUpResponse, error) {
	useCaseRequest := request.getUseCase()

	useCaseResponse, err := instance.topUpUseCase.TopUp(useCaseRequest)
	if err != nil {
		return topUpResponse{}, helper.WrapError(err)
	}

	return topUpResponse{}.fromUseCase(useCaseResponse), nil
}

func (topUpRequest) parse(requestBody string) (request topUpRequest, err error) {
	err = json.Unmarshal([]byte(requestBody), &request)
	if err != nil {
		return topUpRequest{}, helper.WrapError(err)
	}

	return request, nil
}

func (instance topUpRequest) getUseCase() interactor.TopUpRequest {
	return interactor.TopUpRequest{
		UserID:    instance.Params.UserID,
		Increment: instance.Params.Amount,
	}
}

func (topUpResponse) fromUseCase(response interactor.TopUpResponse) topUpResponse {
	return topUpResponse{
		Ok:      response.Ok,
		Message: response.Message,
	}
}
