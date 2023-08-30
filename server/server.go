package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"userapi/container"
	"userapi/handlers"
)

type Server struct {
	httpServer *http.Server
}

func Run(port string, containerInstance container.Container) error {
	e := echo.New()

	routes(e, containerInstance)

	return e.Start(":" + port)
}

func routes(e *echo.Echo, container container.Container) {
	handlers.NewAuthHandler(&container).SetRoutes(e, e.Group("/account"))
	handlers.NewUsersHandler(&container).SetRoutes(e.Group("/users"))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
