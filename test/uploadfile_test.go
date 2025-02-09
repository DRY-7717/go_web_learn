package test

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//go:embed resources/*.gohtml
var uploads embed.FS
var mytemplatesuploads = template.Must(template.ParseFS(uploads, "resources/*.gohtml"))

func UploadFile(w http.ResponseWriter, r *http.Request) {
	mytemplatesuploads.ExecuteTemplate(w, "uploadfile.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// ini proses ambil name dari input file
	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}

	// ini mengatur tempat penyimpanan filenya dimana
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)

	name := r.PostFormValue("name")

	mytemplatesuploads.ExecuteTemplate(w, "uploadsuccess.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

func TestUploadFile(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/uploadfile", UploadFile)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()

}

//go:embed resources/coba.png
var uploadfiletest []byte
func TestUnitUploadFile(t *testing.T) {

	// ini buat testingnya
	requestBody := new(bytes.Buffer)
	writter := multipart.NewWriter(requestBody)
	writter.WriteField("name", "Bima Arya Wicaksana")

	file, _ := writter.CreateFormFile("file", "test.png")
	file.Write(uploadfiletest)
	writter.Close()

	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	request.Header.Set("Content-Type", writter.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}
