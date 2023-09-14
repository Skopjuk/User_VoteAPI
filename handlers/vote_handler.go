package handlers

import (
	"github.com/labstack/echo/v4"
	"userapi/container"
)

type VotesHandler struct {
	container *container.Container
}

func NewVotesHandler(container *container.Container) *VotesHandler {
	return &VotesHandler{container: container}
}

func (v *VotesHandler) SetRoutes(g *echo.Group) {
	g.POST("/", v.Vote)
	g.GET("/", v.GetAllVotes)
}
