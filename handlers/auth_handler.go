package handlers

import (
	"github.com/labstack/echo/v4"
	"userapi/container"
)

type AuthHandler struct {
	container *container.Container
}

func NewAuthHandler(container *container.Container) *AuthHandler {
	return &AuthHandler{container: container}
}

func (a *AuthHandler) SetRoutes(g *echo.Group) {
	g.PUT("/:id", a.ChangePassword)
}
