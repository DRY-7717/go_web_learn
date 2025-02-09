package test

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}
type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}

func (handler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer func() {
		err := recover()
		fmt.Println("RECOVER: ", err)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	handler.Handler.ServeHTTP(w, r)
}

func TestServerHttp(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executed Handler")
		fmt.Fprint(w, "hello world")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executed Handler foo")
		fmt.Fprint(w, "hello foo")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executed Handler panic")
		panic("ups")
	})

	// jadi disini itu logmiddleware bisa masuk ke dalam errorhandler atau error middlerware jadi agar ketika log middleware dijalanin error handler juga berjalan tapi khusus menangkap error aja.
	
	logmiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorhandler := &ErrorHandler{
		Handler: logmiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorhandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
