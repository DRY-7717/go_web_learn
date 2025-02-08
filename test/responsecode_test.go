package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)




func ResponseCode(w http.ResponseWriter, r *http.Request)  {
	
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(400)  // ini untuk menggunakan response code
		fmt.Fprintf(w, "name is empty!")
	}else{
		w.WriteHeader(200)
		fmt.Fprintf(w, "hi, %s", name)
	}

}

func TestResponseCodeSuccess(t *testing.T)  {

	request, _ := http.NewRequest(http.MethodGet,"http://localhost:8080?name=bima", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	
	fmt.Println("Status Code:",response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
func TestResponseCodeFailed(t *testing.T)  {

	request, _ := http.NewRequest(http.MethodGet,"http://localhost:8080?name=", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	
	fmt.Println("Status Code:",response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}