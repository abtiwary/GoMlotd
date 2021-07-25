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

	mServer, err := metalserver.NewServer(
		"localhost",
		8088,
		mlotdDB,
	)

	// Run the server
	mServer.StartHTTPServer()
}
