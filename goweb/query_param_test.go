package belajar_golang_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		// suffix f means with formatter
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Fikri", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryPrameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	if firstName == "" || lastName == "" {
		fmt.Fprint(w, "Hello")
	} else {
		// suffix f means with formatter
		fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
	}
}

func TestMultipleQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?first_name=Fikri&last_name=Fatih", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryPrameter(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultiplePrameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query["name"]

	fmt.Fprintf(w, "Hello %s", strings.Join(name, " "))
}

// use same key param
func TestMultipleParamValues(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Fikri&name=Al&name=Fatih", nil)
	recorder := httptest.NewRecorder()

	MultiplePrameterValues(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}
