package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// create endpoints
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "Hello get")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
