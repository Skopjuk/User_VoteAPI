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

func (a *AuthHandler) SetRoutes(e *echo.Echo, group *echo.Group) {
	g := e.Group("/account", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := a.UserIdentity(c)
			if err != nil {
				return err
			}
			if c.Get("userRole") != "admin" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"error": "you should be admin for this action",
				})
			}
			return next(c)
		}
	})
	g.PUT("/:id/change_password", a.ChangePassword)
	g.PATCH("/:id/update_user", a.UpdateUser)
	g.DELETE("/:id/delete", a.DeleteUser)
}
