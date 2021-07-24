package metalserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abtiwary/gomlotd/metaldb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router  *chi.Mux
	IP      string
	Port    int
	MetalDB *metaldb.MetalDatabase
}

func NewServer(ip string, port int, db *metaldb.MetalDatabase) (*Server, error) {
	nMlotdServer := new(Server)
	nMlotdServer.IP = ip
	nMlotdServer.Port = port

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	nMlotdServer.Router = r

	nMlotdServer.MetalDB = db

	return nMlotdServer, nil
}

func (s *Server) initAPI() {
	s.Router.Get("/api/v1/recommendations", s.HandleGetRecommendations)
}

func (s *Server) StartHTTPServer() *http.Server {
	s.initAPI()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", s.IP, s.Port),
		Handler: s.Router,
	}
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Println("ListenAndServe(): %v", err)
	}

	return srv
}

func (s *Server) HandleGetRecommendations(w http.ResponseWriter, r *http.Request) {
	mrs, err := s.MetalDB.GetRecommendations()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mrs)
}