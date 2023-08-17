package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"userapi/container"
)

type AuthHandler struct {
	container *container.Container
}

func NewAuthHandler(container *container.Container) *AuthHandler {
	return &AuthHandler{container: container}
}

func (a *AuthHandler) SetRoutes(e *echo.Echo) {
	g := e.Group("/auth", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := a.UserIdentity(c)
			if err != nil {
				return err
			}
			if c.Get("userRole") != "admin" {
				c.JSON(http.StatusForbidden, map[string]interface{}{
					"error": "you should be admin for this action",
				})
				return err
			}
			return next(c)
		}
	})
	g.PUT("/:id", a.ChangePassword)
	g.PATCH("/:id", a.UpdateUser)
}
