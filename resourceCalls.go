package main

import (
	"encoding/json"
	"net/http"
)

func getRaces(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	races := getAllRaces()
	response, err := json.Marshal(races)
	if err != nil {
		w.WriteHeader(505)
		writeMessage(w, err.Error())
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func getClasses(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	classes := getAllClasses()
	response, err := json.Marshal(classes)
	if err != nil {
		w.WriteHeader(505)
		writeMessage(w, err.Error())
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
