package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (receiver MyPage) SayHello(name string) string {
	return "Hello " + name + ". My name is is " + receiver.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Fikri",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Fikri",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionMap(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	// create additional function to global function golang
	t = t.Funcs(map[string]interface{}{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "FIKRI",
	})
}

func TestTemplateFunctionMap(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// pass to function pipepines
func TemplateFunctionPipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	// create additional function to global function golang
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(s string) string {
			return "Hello " + s
		},
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	})

	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "FIKRI",
	})
}

func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
