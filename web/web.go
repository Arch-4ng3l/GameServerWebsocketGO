package web

import (
	"net/http"

	"github.com/Arch-4ng3l/GoServerHololens/storage"
)

type WebServer struct {
	store storage.Storage
}

func NewWebServer(store storage.Storage) *WebServer {
	return &WebServer{
		store,
	}
}

func (s *WebServer) Init() {
	fs := http.FileServer(http.Dir("./web/svelte-app/public"))

	http.Handle("/", fs)
}
