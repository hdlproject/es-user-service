package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type (
	HttpServer struct {
		engine    *gin.Engine
		apiRouter *gin.RouterGroup
		port      string
	}
)

var (
	httpServerInstance *HttpServer
)

func GetHTTPServer(port string) *HttpServer {
	if httpServerInstance == nil {
		httpServerInstance = newHTTPServer(port)
	}

	return httpServerInstance
}

func newHTTPServer(port string) *HttpServer {
	engine := gin.Default()
	apiRouter := engine.Group("/api")

	return &HttpServer{
		engine:    engine,
		apiRouter: apiRouter,
		port:      port,
	}
}

func (instance *HttpServer) Serve() {
	instance.engine.Run(fmt.Sprintf(":%s", instance.port))
}
