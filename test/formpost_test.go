package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)



func FormPost(w http.ResponseWriter, r *http.Request)  {
	
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		 //r.PostFormValue("firstname") (sebenernya cara ini lebih simpel, karna tidak harus parsing dan ambil data manual seperti diatas yang parsing dan dibawah yang ambil data secara manual, cara ini langsung menggunakan parsing dan mengambil data.)

		firstname := r.PostForm.Get("firstname")
		lastname := r.PostForm.Get("lastname")

		fmt.Fprintf(w,"Nama saya %s %s", firstname, lastname)


}



func TestFormPost(t *testing.T)  {
	// ini cara test form postnya
	requestbody := strings.NewReader("firstname=Bima&lastname=Wicaksana")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestbody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))




}

