package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code < 499 {
		log.Println("Respomding with err 5XX : ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJOSN(w, code, errResponse{
		Error: msg,
	})
}
func respondWithJOSN(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshall json response: ", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
