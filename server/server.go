package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Arch-4ng3l/GoServerHololens/assets"
	"github.com/Arch-4ng3l/GoServerHololens/storage"
	"github.com/Arch-4ng3l/GoServerHololens/types"
	"github.com/Arch-4ng3l/GoServerHololens/util"
	"github.com/Arch-4ng3l/GoServerHololens/web"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
	store storage.Storage
}

type Conn struct {
	Remote string `json:"remote"`
	Local  string `json:"local"`
}

var HomeDir string

func NewServer(store storage.Storage, homeDir string) *Server {

	HomeDir = homeDir

	return &Server{
		conns: make(map[*websocket.Conn]bool),
		store: store,
	}
}

func (s *Server) Run(addr string) {

	webserver := web.NewWebServer(s.store, HomeDir)
	webserver.Init()

	http.Handle("/api", websocket.Handler(s.handleConns))
	http.HandleFunc("/api/assets", s.handleAssets)
	http.HandleFunc("/api/conns", s.connHandler)
	http.HandleFunc("/api/login", s.handleLogin)
	http.HandleFunc("/api/auth", s.handleAuth)

	fmt.Println("Listening on Address :" + addr)
	http.ListenAndServe(addr, nil)
}

func (s *Server) handleAuth(w http.ResponseWriter, r *http.Request) {

	var tokenObj struct {
		Token string `json:"token"`
	}

	json.NewDecoder(r.Body).Decode(&tokenObj)

	if s.AuthJWT(tokenObj.Token) {
		w.WriteHeader(200)

		return
	}
	w.WriteHeader(400)

}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	user := &types.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusTeapot)
	}
	user2 := s.store.GetUser(user.Username)
	if user2 == nil {
		w.WriteHeader(http.StatusTeapot)

		return
	}

	hashPass := util.CreateHash(user.Password)
	if hashPass != user2.Password || user.Username != user2.Username {
		w.WriteHeader(http.StatusTeapot)

		return
	}

	token, err := util.CreateJWT(user.Username, hashPass)
	if err != nil {

		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (s *Server) connHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		s.handleGetConns(w, r)
	case "POST":
		s.handleCloseConn(w, r)
	}
}

func (s *Server) handleCloseConn(w http.ResponseWriter, r *http.Request) {
	connReq := &Conn{}
	if err := json.NewDecoder(r.Body).Decode(connReq); err != nil {
		util.PrintError(err)

		return
	}

	for conn := range s.conns {
		if conn.RemoteAddr().String() == connReq.Remote {
			w.WriteHeader(200)
			s.closeConn(conn)
		}
	}

}

func (s *Server) handleGetConns(w http.ResponseWriter, r *http.Request) {

	var conns []Conn
	for i := range s.conns {
		conn := Conn{}

		conn.Remote = i.RemoteAddr().String()
		conn.Local = i.LocalAddr().String()
		conns = append(conns, conn)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(conns)
}

func (s *Server) handleAssets(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		s.handleGetAssets(w, r)
	case "POST":
		s.handleCreateAssets(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}
}

func (s *Server) handleCreateAssets(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.PrintError(err)

		return
	}
	defer file.Close()
	writeFile, err := os.Create(HomeDir + "/Hololens/assets/" + header.Filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.PrintError(err)

		return
	}
	if _, err := io.Copy(writeFile, file); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.PrintError(err)

		return
	}
	w.WriteHeader(200)

	assets.Init(HomeDir)

}

func (s *Server) handleGetAssets(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets.zip")
}

func (s *Server) handleConns(conn *websocket.Conn) {

	defer conn.Close()

	s.conns[conn] = true
	str := fmt.Sprintf("New Connection => Connection Number %d", len(s.conns))
	util.PrintLog(str)
	err := s.handleConn(conn)
	if err.Error() != "EOF" {
		util.PrintError(err)
	}
}

func (s *Server) handleConn(conn *websocket.Conn) error {

	var err error

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	//s.initConn(encoder, decoder)

	object := &types.Object{}

	defer s.closeConn(conn)

	for {
		if err = decoder.Decode(object); err != nil {

			return err
		}
		if object.Name == "Player" {
			err = s.handlePlayer(object, encoder)

		} else if object.Name == "Web" {
			err = s.handleWeb(int(object.X), encoder)
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
	util.PrintLog("Connection Closed")
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

	objs := make(chan *types.Object)
	go s.store.GetObjects(objs, player)
	var objects []*types.Object
	for obj := range objs {
		fmt.Println(obj.Name)
		objects = append(objects, obj)
	}

	err := encoder.Encode(objects)

	return err
}

func (s *Server) handleObject(object *types.Object) error {

	fmt.Println(object.Name)

	return s.store.UpdateObject(object)
}

func (s *Server) handleWeb(startID int, encoder *json.Encoder) error {

	objects, err := s.store.GetObjectsWeb(startID)
	if err != nil {
		return err
	}

	return encoder.Encode(objects)
}

func (s *Server) AuthJWT(tokenString string) bool {

	token, err := util.ValidateJWT(tokenString)
	if err != nil {
		return false
	}

	claims := token.Claims.(jwt.MapClaims)

	user2 := s.store.GetUser(claims["username"].(string))

	if claims["username"] != user2.Username || claims["password"] != user2.Password {
		return false
	}

	return true
}
