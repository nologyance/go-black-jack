package controller

import (
	"black-jack/repository"
	"net/http"
)

type PlayerController interface {
	PostPlayer(w http.ResponseWriter, r *http.Request)
}

type playerContoller struct {
	pr repository.PlayerRepository
}

func NewPlayerController(pr repository.PlayerRepository) PlayerController {
	return &playerContoller{pr}
}

func (pc *playerContoller) PostPlayer(w http.ResponseWriter, r *http.Request) {
	// body := make([]byte, r.ContentLength)
	// r.Body.Read(body)
	// var playerRequest dto.PlayerRequest
	// json.Unmarshal(body, &playerRequest)

	// player := entity.PlayerEntity{}

	// // insert

	// w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(1))
	w.WriteHeader(201)
}