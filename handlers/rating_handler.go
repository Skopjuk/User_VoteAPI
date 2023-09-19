package handlers

import (
	"github.com/labstack/echo/v4"
	"userapi/container"
)

type RatingHandler struct {
	container *container.Container
}

func NewRatingHandler(container *container.Container) *RatingHandler {
	return &RatingHandler{container: container}
}

func (r *RatingHandler) SetRoutes(g *echo.Group) {
	g.GET("/:id", r.GetRatingByUserId)
}
