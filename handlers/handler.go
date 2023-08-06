package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"userapi/repositories"
)

type Handler struct {
	db      *sqlx.DB
	logging *logrus.Logger
	router  *repositories.UsersRepository
}

func NewHandler(logging *logrus.Logger, db *sqlx.DB) *Handler {
	return &Handler{
		logging: logging,
		db:      db,
		router:  repositories.NewUsersRepository(db),
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()

	//	auth := router.Group("/auth")
	api := router.Group("/api")
	user := router.Group("/users")

	api.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))
	{
		user.GET("/", h.GetAll)
		user.PUT("/:id", h.UpdateUser)
		user.GET("/:id", h.GetUserById)
		user.GET("/count", h.GerNumberOfUsers)

		user.POST("/", h.SignUp)
		//		auth.POST("/sign-in", h.SignIn)

		api.PUT("/:id", h.ChangePassword)
	}

	return router
}
