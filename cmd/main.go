package main

import (
	"black-jack/controller"
	"black-jack/repository"

	"github.com/labstack/echo"
)

var pr = repository.NewPlayerRepository()
var pc = controller.NewPlayerController(pr)

func main() {
	e := echo.New()
	e.GET("/players", echo.HandlerFunc(pc.GetPlayers))
	e.POST("/players", echo.HandlerFunc(pc.PostPlayer))
	e.Logger.Fatal(e.Start(":8080"))
}
