package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	w.Header().Add("X-Powered-By", "afaf-tech")
	fmt.Fprintf(w, "%v\n", contentType)
}

func TestRequestHeaders(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	poweredBy := recorder.Header().Get("X-Powered-By") // in case sensitive to call

	fmt.Println(string(body))
	fmt.Println(poweredBy)
}
