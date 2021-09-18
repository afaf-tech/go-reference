package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// middleware adaah fitur diman kita bisa menambahkan kode sebelum dan setelah sebuah handler dieksekusi
// sayangnya di go tidak ada konsep middleware, namun karena struktur handler yang baik menggunakan interface, kita bisa
// 		 membuat middleware sendiri menggunakan handler

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before execute Handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("terjadi error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
		fmt.Println("Handler executed")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Foo!")
		fmt.Println("Handler Foo executed")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("ups")
	})

	logmiddleware := &LogMiddleware{
		Handler: mux,
	}
	// errorHandler wii be the top level middleware
	errorHandler := &ErrorHandler{
		Handler: logmiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
