package main

import (
	"black-jack/controller"
	"black-jack/repository"
	"net/http"
)

var pr = repository.NewPlayerRepository()
var pc = controller.NewPlayerController(pr)
var ro = controller.NewRouter(pc)

func main() {
	http.HandleFunc("/player/new", ro.HandlePlayerRequest)
	http.ListenAndServe(":8080", nil)
}
