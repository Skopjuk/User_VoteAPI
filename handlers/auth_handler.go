package handlers

import (
	"github.com/labstack/echo/v4"
	"userapi/server"
)

type AuthHandler struct {
	//	logging *logrus.Logger
	container *server.Container
}

func NewAuthHandler(container *server.Container) *AuthHandler {
	return &AuthHandler{container: container}
}

func (a *AuthHandler) SetRoutes(g *echo.Group) {
	g.PUT("/:id", a.ChangePassword)
}

//
//func (u *AuthHandler) InitRoutes() *echo.Echo {
//	router := authRoutes()
//
//	return router
//}
