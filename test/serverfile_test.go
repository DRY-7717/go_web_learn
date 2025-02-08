package test

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)


func ServerFile(w http.ResponseWriter, r *http.Request)  {
	// jadi servefile itu berguna untuk mengambil file sesuai yang kita inginkan, contohnya kaya tranksasi sukses halaman mau diarahkan kemana atau data yang dicari tidak ditemukan itu bisa menjadi seperti ini.
	
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	}else{
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}


func TestServerFile(t *testing.T)  {
	
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServerFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}

//go:embed resources/ok.html
var responseOk string

//go:embed resources/notfound.html
var responseNotFound string

func ServerFileGoEmbed(w http.ResponseWriter, r *http.Request)  {
	
	// menggunakan go embed itu bisa mempermudah pemanggilan file seperti dibawah tapi kendalanya dia tidak bisa memanggil file yang ada diluar folder
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, responseOk)
	}else{
		fmt.Fprint(w, responseNotFound)
	}
}


func TestServerFileGoEmbed(t *testing.T)  {
	
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServerFileGoEmbed),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}


