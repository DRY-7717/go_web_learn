package test

import (
	"fmt"
	"net/http"
	"testing"
)




func RedirectTo(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "hello redirect")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request)  {
	// logic

	http.Redirect(w,r, "/redirect-to",http.StatusTemporaryRedirect)
}



func TestRedirect(t *testing.T)  {

	mux := http.NewServeMux()


	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-from", RedirectFrom)


	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}


	err := server.ListenAndServe()


	if err != nil {
		panic(err)
	}


}