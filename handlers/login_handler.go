package handlers

import (
	"github.com/labstack/echo/v4"
	"userapi/container"
)

type LoginHandler struct {
	container *container.Container
}

func NewLoginHandler(container *container.Container) *LoginHandler {
	return &LoginHandler{container: container}
}

func (l *LoginHandler) SetRoutes(g *echo.Group) {
	g.POST("/login", l.SignIn)

}
