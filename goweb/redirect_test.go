package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Redirect")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	// whatever logic you want
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}
func RedirectOut(w http.ResponseWriter, r *http.Request) {
	// whatever logic you want
	http.Redirect(w, r, "https://github.com/afaf-tech", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
