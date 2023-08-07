package web

import (
	"net/http"

	"github.com/Arch-4ng3l/GoServerHololens/storage"
)

type WebServer struct {
	store storage.Storage
}

var HomeDir string

func NewWebServer(store storage.Storage, homeDir string) *WebServer {
	HomeDir = homeDir

	return &WebServer{
		store,
	}
}

func (s *WebServer) Init() {
	path := HomeDir + "/Hololens/web/svelte-app/public"
	fs := http.FileServer(http.Dir(path))

	http.Handle("/", fs)
}
