package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Arch-4ng3l/GoServerHololens/storage"
	"github.com/Arch-4ng3l/GoServerHololens/types"
	"github.com/TwiN/go-color"
	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
	store storage.Storage
}

func NewServer(store storage.Storage) *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
		store: store,
	}
}

func (s *Server) Run(addr string) {

	http.Handle("/", websocket.Handler(s.handleConns))
	http.HandleFunc("/assets", s.handleAssets)

	fmt.Println("Listening on Address :" + addr)
	http.ListenAndServe(addr, nil)
}

func (s *Server) handleAssets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleGetAssets(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (s *Server) handleGetAssets(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets.zip")
}

func (s *Server) handleConns(conn *websocket.Conn) {
	defer conn.Close()

	s.conns[conn] = true
	fmt.Printf("%s:%s%s [+] New Connection%s => Connection Number %d\n", time.Now().Format(time.Kitchen), color.Green, color.Bold, color.Reset, len(s.conns))
	err := s.handleConn(conn)
	log.Println(time.Now().Format(time.Kitchen) + " " + color.Red + color.Bold + err.Error() + color.Reset)
}

func (s *Server) handleConn(conn *websocket.Conn) error {
	var err error

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	s.initConn(encoder, decoder)

	object := &types.Object{}

	defer s.closeConn(conn)

	for {
		if err = decoder.Decode(object); err != nil {
			return err
		}
		fmt.Println(object)
		if object.Name == "Player" {
			fmt.Println("Got Player")
			err = s.handlePlayer(object, encoder)

		} else {
			err = s.handleObject(object)
		}

		if err != nil {
			return err
		}

	}

}

func (s *Server) closeConn(conn *websocket.Conn) {
	conn.Close()
	delete(s.conns, conn)
	fmt.Println("[-] Connection Closed")
}

func (s *Server) initConn(encoder *json.Encoder, decoder *json.Decoder) error {

	params := &types.Setting{}

	if err := decoder.Decode(params); err != nil {
		return err
	}

	types.RenderDistance = params.RenderDistance

	return nil
}

func (s *Server) handlePlayer(player *types.Object, encoder *json.Encoder) error {
	objects, err := s.store.GetObjects(player)
	if err != nil {
		return err
	}

	err = encoder.Encode(objects)

	return err
}

func (s *Server) handleObject(object *types.Object) error {

	return s.store.UpdateObject(object)
}