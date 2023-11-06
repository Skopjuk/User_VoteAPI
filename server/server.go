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

func Run(port string, redisPort string, containerInstance container.Container) error {
	e := echo.New()

	routes(e, containerInstance)

	return e.Start(":" + port)
}

func routes(e *echo.Echo, container container.Container) {
	handlers.NewAccountHandler(&container).SetRoutes(e.Group("/account", handlers.UserIdentityMiddleware))
	handlers.NewUsersHandler(&container).SetRoutes(e.Group("/users"))
	handlers.NewLoginHandler(&container).SetRoutes(e.Group("/authentication"))
	handlers.NewVotesHandler(&container).SetRoutes(e.Group("/vote", handlers.UserIdentityMiddleware))
	handlers.NewRatingHandler(&container).SetRoutes(e.Group("/rating"))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
