package api

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/input_port"
)

type (
	registerRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func (registerRequest) parse(ctx *gin.Context) (request registerRequest, err error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return registerRequest{}, helper.WrapError(err)
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return registerRequest{}, helper.WrapError(err)
	}

	return request, nil
}

func (instance registerRequest) getUseCase() input_port.RegisterRequest {
	return input_port.RegisterRequest{
		Username: instance.Username,
		Password: instance.Password,
	}
}
