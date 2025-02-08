package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {

	textHtml := `<html> <body> <h1>{{.}}</h1> </body> </html>`
	// disini kalo namanya simple maka templatenya akan bernama simple
	// t, err := template.New("SIMPLE").Parse(textHtml)
	// if err != nil {
	// 	panic(err)
	// }

	// bia juga dengan cara seperti ini, ini bedanya errornya udah dihandle
	t := template.Must(template.New("SIMPLE").Parse(textHtml))
	t.ExecuteTemplate(w, "SIMPLE", "hello bima")

}

func TestSimpleHtml(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

func HTMLFile(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./resources/index.gohtml"))
	t.ExecuteTemplate(w, "index.gohtml", "hello bima")

}

func TestHTMLFile(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

func HTMLFileDirectory(w http.ResponseWriter, r *http.Request) {
	// ini untuk menambil semua file yang bertipe .gohtml
	t := template.Must(template.ParseGlob("./resources/*gohtml"))
	t.ExecuteTemplate(w, "index.gohtml", "hello bima")

}

func TestHTMLFileDirectory(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFileDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

//go:embed resources/*.gohtml
var files embed.FS

func HTMLFileGoEmbed(w http.ResponseWriter, r *http.Request) {
	// ini untuk menambil semua file yang bertipe .gohtml
	t := template.Must(template.ParseFS(files, "resources/*.gohtml"))
	t.ExecuteTemplate(w, "index.gohtml", "hello embed")

}

func TestHTMLFileGoEmbed(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFileGoEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

func HTMLFileDataMap(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./resources/data.gohtml"))
	t.ExecuteTemplate(w, "data.gohtml", map[string]interface{}{
		"Name":    "Bima Arya Wicaksana",
		"Address": "Depok",
	})

}

func TestHTMLFileDataMap(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFileDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

type Hobbies struct {
	Hobby string
}

type Person struct {
	Name, Address interface{}
	Hobbies       Hobbies
}

func HTMLFileDataStruct(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./resources/data.gohtml"))
	t.ExecuteTemplate(w, "data.gohtml", Person{
		Name:    "Bima Arya Wicaksana",
		Address: "Depok, Jawa Barat",
		Hobbies: Hobbies{
			Hobby: "Coding",
		},
	})

}

func TestHTMLFileDataStruct(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFileDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

func HTMLFileSimpleIf(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./resources/ifcondition.gohtml"))
	t.ExecuteTemplate(w, "ifcondition.gohtml", Person{
		Name:    false,
		Address: "Depok, Jawa Barat",
	})

}

func TestHTMLFileSimpleIf(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFileSimpleIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	// ini cara layouting seperti ini
	t := template.Must(template.ParseFiles(
		"./resources/header.gohtml", "./resources/layout.gohtml", "./resources/footer.gohtml"))

	t.ExecuteTemplate(w, "layout", Person{
		Name:    "Bima Arya Wicaksana",
		Address: "Depok, Jawa Barat",
	})
}

func TestHTMLFileLayout(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

type Introduce struct {
	Name string
}

func (intro Introduce) SayHello(name string) string {

	return "halo, " + name + " nama saya " + intro.Name
}

func HTMLFunc(w http.ResponseWriter, r *http.Request) {
	// didalam template juga bisa kirim function

	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Jhon" }}`))

	t.ExecuteTemplate(w, "FUNCTION", Introduce{
		Name: "Bima",
	})

}

func TestHTMLFunc(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFunc(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

func HTMLFuncGlobal(w http.ResponseWriter, r *http.Request) {
	// didalam template juga bisa kirim function
	// disini juga ada fitur kaya length di javascript untuk menghitung jumlah data atau kata

	t := template.Must(template.New("FUNCTION").Parse(`{{ len .Name }}`))

	t.ExecuteTemplate(w, "FUNCTION", Introduce{
		Name: "Bima",
	})

}

func TestHTMLFuncGlobal(t *testing.T) {

	// begini cara testnya
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFuncGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

func HTMLFuncCustomGlobal(w http.ResponseWriter, r *http.Request) {
	// membuat custom global function

	t := template.New("FUNCTION")
	// ini membuat custom global func nya dulu, kalo ga dibikin itu bakal error
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name }}`))

	t.ExecuteTemplate(w, "FUNCTION", Introduce{
		Name: "Bima",
	})

}
func TestHTMLFuncCustomGlobal(t *testing.T) {
	
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFuncCustomGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}
func HTMLFuncPipeline(w http.ResponseWriter, r *http.Request) {
	// membuat custom global function
	// pipeline function itu bisa meneruskan hasil dari function satu ke fuction yang lain atau juga bisa menggunakan 2 function sekaliguts

	t := template.New("FUNCTION")
	// ini membuat custom global func nya dulu, kalo ga dibikin itu bakal error
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
		"SayHello" : func (value string) string {
			return "Hello, " + value
		},
	})
	// jadi si value .Name ini dari function SayHello akan diteruskan ke function Upper dan resultnya nanti berhakhir di function upper
	t = template.Must(t.Parse(`{{ SayHello .Name | upper  }}`))

	t.ExecuteTemplate(w, "FUNCTION", Introduce{
		Name: "Bima",
	})

}
func TestHTMLFuncPipeline(t *testing.T) {
	
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HTMLFuncPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}
