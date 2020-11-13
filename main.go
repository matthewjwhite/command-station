package main

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/matthewjwhite/command-station/command"
	"github.com/matthewjwhite/command-station/config"
	"github.com/matthewjwhite/command-station/render"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	expectedArgs    = 2
	commandEndpoint = "command"
)

func main() {
	if len(os.Args) != expectedArgs {
		log.Fatalf("No command file path provided")
	}

	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open command file: %v", err)
	}

	config, err := config.Parse(file)
	if err != nil {
		log.Fatalf("Failed to parse command file: %v", err)
	}
	file.Close()

	launchServer(config)
}

func launchServer(config config.Config) {
	var handler http.Handler

	handler = mux.NewRouter()

	// Base command station.
	handler.(*mux.Router).HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		data, err := render.Station(config, commandEndpoint)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(data)
	})

	// Command endpoint.
	handler.(*mux.Router).HandleFunc("/"+commandEndpoint+"/{command}", func(w http.ResponseWriter, r *http.Request) {
		cmd, err := command.Collection(config.Commands).Get(mux.Vars(r)["command"])
		if errors.Is(err, command.ErrUnknown) {
			http.Error(w, "Command does not exist!", http.StatusInternalServerError)
		}

		out, err := cmd.Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(out)
	})

	if config.Timeout != 0 {
		handler = http.TimeoutHandler(handler, time.Second*time.Duration(config.Timeout), "Operation timed out")
	}

	log.Fatal(http.ListenAndServe(":8000", handler))
}
