package handlers

import (
	"github.com/labstack/echo/v4"
	"userapi/container"
)

type UsersHandler struct {
	container *container.Container
}

func NewUsersHandler(container *container.Container) *UsersHandler {
	return &UsersHandler{container: container}
}

func (u *UsersHandler) SetRoutes(g *echo.Group) {
	g.GET("/", u.GetAll)
	g.GET("/:id", u.GetUserById)
	g.GET("/count", u.GetNumberOfUsers)
	g.POST("/", u.SignUp)
}
