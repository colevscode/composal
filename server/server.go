package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/colevscode/composal/player"
	"github.com/colevscode/composal/track"
	"github.com/gorilla/mux"
)

type Payload struct {
	Tracks []track.Track `json:"tracks"`
}

func Play(w http.ResponseWriter, r *http.Request) {
	var payload Payload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, `Bad Request. No body.`, 400)
		return
	}

	for _, t := range payload.Tracks {
		player.AddTrack(&t)
	}

	player.Play(0, w)

	// Mirror back json just for debugging
	// js, err := json.Marshal(payload)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"status": "OK"}`))
}

func RunServer(port int, debug bool) {
	address := fmt.Sprintf(":%d", port)
	if debug {
		fmt.Fprintln(os.Stderr, "Opening server on", address)
	}

	router := mux.NewRouter()
	router.HandleFunc("/play", Play).Methods("POST")
	http.ListenAndServe(address, router)
}
