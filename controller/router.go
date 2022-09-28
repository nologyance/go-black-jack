package controller

import "net/http"

type Router interface {
    HandlePlayerRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
    pc PlayerController
}

func NewRouter(pc PlayerController) Router {
    return &router{pc}
}

func (ro *router) HandlePlayerRequest(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        w.WriteHeader(200)
    case "POST":
        ro.pc.PostPlayer(w, r)
    default:
        w.WriteHeader(405)
    }
}