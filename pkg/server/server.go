package server

import (
	"net/http"

	"github.com/Muha113/task-service/pkg/routers"
	"github.com/Muha113/task-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ListenAddr string
	Config     *Config
	Routes     *gin.Engine
}

func NewServer(cfgPath string) (*Server, error) {
	server := &Server{
		Config: &Config{},
		Routes: &gin.Engine{},
	}

	var err error

	server.Config, err = InitConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	server.Routes = routers.SetupRouter()
	server.ListenAddr = utils.BuildAddrString(server.Config.HttpAddr, server.Config.HttpPort)

	return server, nil
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.ListenAddr, s.Routes)
}
