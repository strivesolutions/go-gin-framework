package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/strivesolutions/go-gin-framework/pkg/api"
	"github.com/strivesolutions/go-gin-framework/pkg/middleware"
	"github.com/strivesolutions/logger-go/pkg/logging"
)

var AuthMiddleware = middleware.Auth

type Server struct {
	Engine *gin.Engine
}

type Options struct {
	NoTrustFundMiddleware bool
}

func CreateServer(options Options) Server {
	server := Server{}
	server.Init(options)

	return server
}

func (s *Server) Init(options Options) {
	if s.Engine == nil {
		s.Engine = gin.Default()

		if !options.NoTrustFundMiddleware {
			s.AddMiddleware(middleware.TrustFundId)
		}
	}
}

func (s *Server) AddMiddleware(middleware gin.HandlerFunc) {
	s.Engine.Use(middleware)
}

func (s *Server) AddRoutes(routes []api.ApiRoute) {
	for _, route := range routes {
		s.AddRoute(route)
	}
}

func (s *Server) AddRoute(route api.ApiRoute) {
	var handlers []gin.HandlerFunc

	if route.Anonymous {
		handlers = make([]gin.HandlerFunc, 1)
		handlers[0] = route.Handler
	} else {
		handlers = make([]gin.HandlerFunc, 2)
		handlers[0] = route.Handler
		handlers[1] = AuthMiddleware
	}

	switch route.MethodType {
	case api.GET:
		s.Engine.GET(route.Path, handlers...)

	case api.POST:
		s.Engine.POST(route.Path, handlers...)

	case api.PUT:
		s.Engine.PUT(route.Path, handlers...)

	case api.DELETE:
		s.Engine.DELETE(route.Path, handlers...)

	}
}

func (s *Server) Start(port int) {
	logging.Info("Starting Server")
	s.Engine.Run(fmt.Sprintf(":%d", port))
}
