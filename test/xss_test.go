package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"html/template" // ini harus html/template yak
)

// golang itu aman dari xss scripting
// ? template.HTML ini yang digunakan untuk merender tag-tag html saya, dalam artian mematikan autoescape dan mengizinkan tag html saja yang berjalan
// ? template.JS ini yang digunakan untuk merender tag-tag html saya, dalam artian mematikan autoescape dan mengizinkan tag html saja yang berjalan
// ? template.CSS ini yang digunakan untuk merender tag-tag html saya, dalam artian mematikan autoescape dan mengizinkan tag html saja yang berjalan

//go:embed resources/*.gohtml
var escapes embed.FS
var mytemplatesescapes = template.Must(template.ParseFS(escapes, "resources/*.gohtml"))

func HTMLFileEscape(w http.ResponseWriter, r *http.Request)  {
	
	mytemplatesescapes.ExecuteTemplate(w, "autoescape.gohtml",map[string]interface{}{
		"Title" : "Test Auto Escape HTML",
		"Body" : "<p> hello world </p>",
	} )
}



func TestHTMLFileEscape(t *testing.T) {
	
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFileEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}


func HTMLFileRemoveEscapeTemplateHtml(w http.ResponseWriter, r *http.Request)  {
	
	mytemplatesescapes.ExecuteTemplate(w, "autoescape.gohtml",map[string]interface{}{
		"Title" : "Test Auto Escape HTML",
		"Body" : template.HTML(r.URL.Query().Get("body")),
	} )
}


func TestHTMLFileRemoveEscapeTemplateHtml(t *testing.T) {
	
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFileRemoveEscapeTemplateHtml(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}


func TestServerEscape(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(HTMLFileRemoveEscapeTemplateHtml),
	}

	server.ListenAndServe()
}

