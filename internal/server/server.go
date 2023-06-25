package server

import (
	"FT_ServerClient/pkg/tools"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const address = "localhost:8000"

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	var mux *http.ServeMux
	return &Server{mux}
}

func (s *Server) initHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", s.hello)
	mux.HandleFunc("/hello_username", s.helloUsername)
	return mux
}

func (s *Server) Start() {
	s.mux = s.initHandlers()
	log.Fatal(http.ListenAndServe(address, s.mux))
}

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request /hello")

	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(map[string]string{
		"message": "hello, stranger!",
	})

	if tools.MinorError(err, "Failed to marshal response") {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	_, err = w.Write(resp)
	tools.MinorError(err, "Failed to write response")
}

func (s *Server) helloUsername(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request /hello_username")

	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	var requestData map[string]string

	err := json.NewDecoder(r.Body).Decode(&requestData)
	tools.MinorError(err, "Failed to decode request body")

	tools.MinorError(err, "Failed to unmarshal request body")

	msg := fmt.Sprintf("hello, %s!", requestData["username"])

	resp, err := json.Marshal(map[string]string{
		"message": msg,
	})

	if tools.MinorError(err, "Failed to marshal response") {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	_, err = w.Write(resp)
	tools.MinorError(err, "Failed to write response")
}
