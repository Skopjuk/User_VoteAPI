package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	api := router.Group("/api")
	api.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		{
			api.GET("/all_users", h.GetAll)
			api.PUT("/:id", h.UpdateUser)
			api.GET("/:id", h.GetUserById)
		}
	}

	return router
}
