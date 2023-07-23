package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	db      *sqlx.DB
	logging *logrus.Logger
}

func NewHandler(logging *logrus.Logger, db *sqlx.DB) *Handler {
	return &Handler{
		logging: logging,
		db:      db,
	}
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
