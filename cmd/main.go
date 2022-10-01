package main

import (
	"black-jack/controller"
	"black-jack/repository"
	"log"
	"net/http"
)

var pr = repository.NewPlayerRepository()
var pc = controller.NewPlayerController(pr)

func main() {
	http.Handle("/player", http.HandlerFunc(pc.HandlePlayerRequest))
	err := http.ListenAndServe(":8080", nil)
	if (err != nil) {
		log.Fatal(err)
	}
}
