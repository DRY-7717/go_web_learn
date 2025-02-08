package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie) // ini cara membuat cookie ada menginisialisasi cookie di golang

	cookie.Name = "Bima-Cookie"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"         // path ini tempat untuk menaruh cookie mau dimana kalo pathnya hanya "/" itu berarti cookie ini bisa di tarus ke seluruh halaman website tapi kalo mau satu saja tambahkan endpoint setelah pathnya
	http.SetCookie(w, cookie) // ini cara menset atau memasang cookie di browser
	fmt.Fprintln(w, "Berhasil memasang cookie")

}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Bima-Cookie")
	if err != nil {
		fmt.Fprintf(w, "has no cookie!")
	} else {
		fmt.Fprintf(w, "hi, %s!", cookie.Value)
	}
}

func TestRunningCookie(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestSetCookie(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080?name=Bima", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)
	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {

		fmt.Printf("%s : %s \n", cookie.Name, cookie.Value)
	}

}
func TestGetCookie(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080?name=wicaksana", nil)
	cookie := new(http.Cookie) 
	cookie.Name = "Bima-Cookie"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/" 
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
