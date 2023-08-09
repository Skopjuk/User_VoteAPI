package handlers

import (
	"github.com/labstack/echo/v4"
	"userapi/server"
)

type UsersHandler struct {
	container *server.Container
}

func NewUsersHandler(container *server.Container) *UsersHandler {
	return &UsersHandler{container: container}
}

func (u *UsersHandler) SetRoutes(g *echo.Group) {
	g.GET("/", u.GetAll)
	g.PUT("/:id", u.UpdateUser)
	g.GET("/:id", u.GetUserById)
	g.GET("/count", u.GerNumberOfUsers)
	g.POST("/", u.SignUp)

}

//func (u *UsersHandler) InitRoutes() *echo.Echo {
//	router := routes()
//
//	return router
//}
