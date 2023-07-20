package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logging *logrus.Logger
}

func NewHandler(logging *logrus.Logger) *Handler {
	return &Handler{logging: logging}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("sign-in", h.SignIn)
	}
	//	api := router.Group("/api")
	//	{
	////		user := api.Group("/user")
	//		{
	////			user.POST("/:id", h.updateUser)
	//		}
	//	}

	return router
}
