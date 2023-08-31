package bug

import (
	"fmt"
	"net/http"
)

func Server() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/subpath/", HandleSubPath)
	mux.HandleFunc("/", HandleMain)
	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func HandleSubPath(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Got method %s", r.Method)))
}
