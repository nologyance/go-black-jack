package controller

import (
	"black-jack/entity"
	"black-jack/repository"
	"net/http"
)

type PlayerController interface {
	HandlePlayerRequest(w http.ResponseWriter, r *http.Request)
}

type playerContoller struct {
	pr repository.PlayerRepository
}

func NewPlayerController(pr repository.PlayerRepository) PlayerController {
	return &playerContoller{pr}
}

func (pc *playerContoller) HandlePlayerRequest(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
		pc.pr.InsertPlayer(entity.PlayerEntity{})
        w.WriteHeader(200)
    case http.MethodPost:
        pc.PostPlayer(w, r)
    default:
        w.WriteHeader(405)
    }
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