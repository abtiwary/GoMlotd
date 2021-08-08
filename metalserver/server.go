package metalserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abtiwary/gomlotd/metaldb"
	mlotd "github.com/abtiwary/gomlotd/metallotd"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Router  *chi.Mux
	IP      string
	Port    int
	Srv     *http.Server
	MetalDB *metaldb.MetalDatabase
	Exit    chan struct{}
}

func NewServer(ip string, port int, db *metaldb.MetalDatabase) (*Server, error) {
	nMlotdServer := new(Server)
	nMlotdServer.IP = ip
	nMlotdServer.Port = port

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	nMlotdServer.Router = r

	nMlotdServer.MetalDB = db
	nMlotdServer.Exit = make(chan struct{})

	return nMlotdServer, nil
}

func (s *Server) initAPI() {
	s.Router.Get("/api/v1/recommendations", s.HandleGetRecommendations)
	s.Router.Post("/api/v1/recommendation", s.HandleSetRecommendation)
}

func (s *Server) StartHTTPServer() *http.Server {
	s.initAPI()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", s.IP, s.Port),
		Handler: s.Router,
	}

	go func(exitChan chan struct{}) {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("ListenAndServe(): %v", err)
		}
	}(s.Exit)

	s.Srv = srv
	return srv
}

func (s *Server) StopHTTPServer() {
	log.Debug("stopping HTTP server")
	defer s.Srv.Close()
	close(s.Exit)
}

func (s *Server) HandleGetRecommendations(w http.ResponseWriter, r *http.Request) {
	mrs, err := s.MetalDB.GetRecommendations()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if len(mrs) == 0 || mrs == nil {
		emptyMrs := make([]string, 0)
		json.NewEncoder(w).Encode(emptyMrs)
	} else {
		json.NewEncoder(w).Encode(mrs)
	}

}

func (s *Server) HandleSetRecommendation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	rec := struct {
		Video string `json:"video"`
	}{}
	err := decoder.Decode(&rec)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("Got video: %v\n", rec.Video)
	if rec.Video == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	metalLoTD := mlotd.NewMetalLinkOfTheDay(rec.Video)
	err = metalLoTD.GetDetails()
	if err != nil {
		log.WithError(err).Info("could not get video details")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.WithField("title", metalLoTD.VideoTitle).Debug("title of video")

	mr := metaldb.MetalRecommendation{}
	mr.URL = metalLoTD.URL
	mr.VideoID = metalLoTD.VideoID
	mr.VideoTitle = metalLoTD.VideoTitle

	err = s.MetalDB.StoreRecommendation(&mr)
	if err != nil {
		log.WithError(err).Info("error writing metal recommendation to the database")
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

}
