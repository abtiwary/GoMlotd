package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/abtiwary/gomlotd/metaldb"
	"github.com/abtiwary/gomlotd/metalserver"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	endCh := make(chan os.Signal, 1)
	signal.Notify(endCh, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

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
	defer mServer.StopHTTPServer()

	// Run the server
	mServer.StartHTTPServer()

	for {
		select {
		case <-endCh:
			log.Debug("terminating...")
			os.Exit(-1)
		}
	}
}
