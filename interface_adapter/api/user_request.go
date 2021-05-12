package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/interactor"
	"io/ioutil"
)

type (
	registerRequest struct{}
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

func (instance registerRequest) getUseCase() interactor.RegisterRequest {
	return interactor.RegisterRequest{}
}
