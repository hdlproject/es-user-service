package api

import "github.com/hdlproject/es-user-service/use_case/interactor"

type (
	registerResponse struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
		UserID  uint   `json:"user_id"`
	}
)

func (registerResponse) fromUseCase(request interactor.RegisterResponse) registerResponse {
	return registerResponse{
		Ok:      request.Ok,
		Message: request.Message,
		UserID:  request.UserID,
	}
}
