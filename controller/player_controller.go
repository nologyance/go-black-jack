package controller

import (
	"black-jack/entity"
	"black-jack/repository"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type PlayerController interface {
	GetPlayers(c echo.Context) error
	PostPlayer(c echo.Context) error
}

type playerContoller struct {
	pr repository.PlayerRepository
}

func NewPlayerController(pr repository.PlayerRepository) PlayerController {
	return &playerContoller{pr}
}

func (pc *playerContoller) GetPlayers(c echo.Context) error {
	return c.JSON(http.StatusOK, pc.pr.SelectPlayers())
}

func (pc *playerContoller) PostPlayer(c echo.Context) error {
	player := new(entity.PlayerEntity)
	if err := c.Bind(player); err != nil {
		log.Fatal(err)
		return err
	}
	pc.pr.InsertPlayer(*player)
	return c.JSON(http.StatusCreated, player)
}