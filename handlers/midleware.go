package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := userIdentity(c)
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
}
