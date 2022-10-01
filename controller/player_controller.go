package controller

import (
	"black-jack/entity"
	"black-jack/repository"
	"encoding/json"
	"net/http"
	"strconv"
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
	player := pc.decode(r)
	pc.pr.InsertPlayer(player)
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(player.Id))
	w.WriteHeader(201)
}

func (pc *playerContoller) decode(r *http.Request) entity.PlayerEntity {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var playerRequest PlayerRequest
	json.Unmarshal(body, &playerRequest)
	return entity.PlayerEntity(playerRequest)
}