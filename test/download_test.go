package test

import (
	"fmt"
	"net/http"
	"testing"
)


func DownloadFile(w http.ResponseWriter, r *http.Request)  {

	filename := r.URL.Query().Get("file")

	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return 
	}

	// ini fungsi untuk downloadnya
	w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	http.ServeFile(w, r, "./resources/" + filename)
	
}



func TestDownloadFile(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}
	server.ListenAndServe()
}