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

//func (s *Server) Run(port string, handler http.Handler) error {
//	s.httpServer = &http.Server{
//		Addr:           ":" + port,
//		MaxHeaderBytes: 1 << 20,
//		Handler:        handler,
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//	}
//
//	logrus.Infof("listen and serve on port %s", port)
//
//	return s.httpServer.ListenAndServe()
//}

func Run(port string, containerInstance container.Container) error {
	e := echo.New()

	routes(e, containerInstance)

	return e.Start(":" + port)
}

func routes(e *echo.Echo, container container.Container) {
	handlers.NewAuthHandler(&container).SetRoutes(e.Group("/auth"))
	handlers.NewUsersHandler(&container).SetRoutes(e.Group("/users"))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
