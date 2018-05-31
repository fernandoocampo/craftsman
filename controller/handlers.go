package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func findByID(w http.ResponseWriter, r *http.Request) {

}

func create(w http.ResponseWriter, r *http.Request) {

}

func update(w http.ResponseWriter, r *http.Request) {

}

func updateState(w http.ResponseWriter, r *http.Request) {

}

func health(w http.ResponseWriter, r *http.Request) {
	// If all is ok, marshal into JSON, write headers and content
	respondWithJSON(w, http.StatusOK, health)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	w.WriteHeader(code)
	w.Write(response)
}
