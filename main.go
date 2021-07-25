package main

import (
	"os"

	"github.com/abtiwary/gomlotd/metaldb"
	"github.com/abtiwary/gomlotd/metalserver"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	mlotdDB, err := metaldb.NewMetalDatabase(
		"localhost",
		5432,
		"mlotd",
		"postgres",
		"postgres",
	)
	if err != nil {
		panic("error creating a database connection")
	}
	defer mlotdDB.DB.Close()

	/*
		metalLoTD := mlotd.NewMetalLinkOfTheDay("https://www.youtube.com/watch?v=kdGlZAWWghY&t=134s")
		err = metalLoTD.GetDetails()
		if err != nil {
			log.WithError(err).Info("could not get video details")
		}
	*/

	//log.WithField("title", metalLoTD.VideoTitle).Debug("title of video")

	/*
		mr := metaldb.MetalRecommendation{}
		mr.URL = metalLoTD.URL
		mr.VideoID = metalLoTD.VideoID
		mr.VideoTitle = metalLoTD.VideoTitle

		err = mlotdDB.StoreRecommendation(&mr)
		if err != nil {
			log.WithError(err).Info("error writing metal recommendation to the database")
		}
	*/

	/*
		mrs, err := mlotdDB.GetRecommendations()
		for _, mr := range mrs {
			fmt.Println(mr)
		}
	*/

	mServer, err := metalserver.NewServer(
		"localhost",
		8088,
		mlotdDB,
	)

	// Run the server
	mServer.StartHTTPServer()

}
