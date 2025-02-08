package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {

	server := http.Server{
		Addr: "localhost:9090",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	}

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
func TestMux(t *testing.T) {

	// mux ini kalo istilah di laravelnya dia ini route
	// mux juga bisa jadi handler
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	// /hi/ kalo penulisannya ujungnya ada / nya itu apapun yang diketikan setelah / akan mengakses ke halaman /hi
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hi juga")
	})
	mux.HandleFunc("/hi/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello juga")
	})

	server := http.Server{
		Addr:    "localhost:7717",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}

func TestRequest(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// ini request bisa mengambil method url, uri dll...
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.URL)
	}

	server := http.Server{
		Addr:    "localhost:7717",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
