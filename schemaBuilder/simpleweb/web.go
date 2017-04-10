package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

type MyServer struct {
	r *mux.Router
}

func main() {

	// Server the static content
	rcommon := mux.NewRouter()
	rcommon.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", &MyServer{rcommon})

	// Answer the forms call back here and process the results in JSON-LD and present back to the user

	// Start the server...
	log.Printf("About to listen on 19900. Go to http://127.0.0.1:19900/")

	err := http.ListenAndServe(":19900", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	s.r.ServeHTTP(rw, req)
}
