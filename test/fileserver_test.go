package test

import (
	"embed"
	_ "embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {

	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()

	// kalo cuman gini aja ini ga akan tampil karna si server memanggil seperti ini: /resources/static/namafile, sedankan kita itu arahkan pathnya langsung ke resources

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

//go:embed resources
var resources embed.FS

func TestFileServerGoEmbed(t *testing.T) {

	// kalo embed itu harus pakai http.FS di http.fileservernya, dan kalo hanya gini saja itu filenya akan balik ke awal ke path seperti ini resources/static/namafile untuk menghindari itu bisa menggunakan fs.Sub

	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
