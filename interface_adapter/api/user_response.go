package api

import (
	"github.com/hdlproject/es-user-service/use_case/input_port"
)

type (
	registerResponse struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
		UserID  uint   `json:"user_id"`
	}
)

func (registerResponse) fromUseCase(request input_port.RegisterResponse) registerResponse {
	return registerResponse{
		Ok:      request.Ok,
		Message: request.Message,
		UserID:  request.UserID,
	}
}
