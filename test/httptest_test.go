package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// http test itu untuk test http, kalo test sebelumnya itu kan kita bukan ngetest tapi bener-bener menjalankan server

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil) // ini itu request mau diakses ke url mana dan parameter ketiga itu body
	recorder := httptest.NewRecorder()                                                 // ini itu kalo di handler ini itu jadi writter
	HelloHandler(recorder, request)

	response := recorder.Result()        // ini untuk mengambil responsenya
	body, _ := io.ReadAll(response.Body) // ini itu untuk membaca response body yang tipenya byte

	fmt.Println(string(body))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // ini untuk mengambil query diurl
	if name == "" {
		fmt.Fprintln(w, "hello")
	} else {
		fmt.Fprintf(w, "hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=nisa", nil) // ini itu request mau diakses ke url mana dan parameter ketiga itu body
	recorder := httptest.NewRecorder()                                                           // ini itu kalo di handler ini itu jadi writter
	SayHello(recorder, request)

	response := recorder.Result()        // ini untuk mengambil responsenya
	body, _ := io.ReadAll(response.Body) // ini itu untuk membaca response body yang tipenya byte

	fmt.Println(string(body))
}

func MultipleQuery(w http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("firstname")
	lastname := r.URL.Query().Get("lastname")
	fmt.Fprintf(w, "halo nama saya %s %s", firstname, lastname)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?firstname=BimaArya&lastname=Wicaksana", nil) // ini itu request mau diakses ke url mana dan parameter ketiga itu body
	recorder := httptest.NewRecorder()                                                                                       // ini itu kalo di handler ini itu jadi writter
	MultipleQuery(recorder, request)

	response := recorder.Result()        // ini untuk mengambil responsenya
	body, _ := io.ReadAll(response.Body) // ini itu untuk membaca response body yang tipenya byte

	fmt.Println(string(body))
}

func MultipleValueQuery(w http.ResponseWriter, r *http.Request) {
	// untuk mendapatkan value dari nama satu query parameter bisa lakukan ini:
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprintln(w, strings.Join(names, ", "))
}

func TestMultipleValueQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Bima&name=Arya&name=Wicaksana", nil) // ini itu request mau diakses ke url mana dan parameter ketiga itu body
	recorder := httptest.NewRecorder()                                                                                    // ini itu kalo di handler ini itu jadi writter
	MultipleValueQuery(recorder, request)

	response := recorder.Result()        // ini untuk mengambil responsenya
	body, _ := io.ReadAll(response.Body) // ini itu untuk membaca response body yang tipenya byte

	fmt.Println(string(body))
}

// header

func GetHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprintln(w, contentType)
}

func TestGetheader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
    request.Header.Add("Content-type", "applicaiton/json")
	recorder := httptest.NewRecorder()
	GetHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}


func SetHeader(w http.ResponseWriter, r *http.Request) {
    // begini cara membuat response header kita sendiri
	w.Header().Add("X-Powered-By", "Powered: Bima Arya Wicaksana") // dia ini gabisa dimasukin ke variable
	fmt.Fprintln(w, "Ok")
}

func TestSetHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
    
	recorder := httptest.NewRecorder()
	SetHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

    fmt.Println(recorder.Header().Get("x-powered-by"))
}

