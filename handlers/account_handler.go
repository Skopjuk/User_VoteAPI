package handlers

import (
	"github.com/labstack/echo/v4"
	"userapi/container"
)

type AccountHandler struct {
	container *container.Container
}

func NewAccountHandler(container *container.Container) *AccountHandler {
	return &AccountHandler{container: container}
}

func (a *AccountHandler) SetRoutes(g *echo.Group) {
	g.PUT("/:id/change_password", a.ChangePassword)
	g.PATCH("/:id/update_user", a.UpdateUser)
	g.DELETE("/:id/delete", a.DeleteUser)
}
