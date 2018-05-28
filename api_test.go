package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
   "fmt"
)

func init(){
	http.HandleFunc("/", showImage)
}

func TestCorrectResponse(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rw := httptest.NewRecorder()
    http.DefaultServeMux.ServeHTTP(rw, req)
    if rw.Code != 200 {            
    t.Fatalf("Expected 200 response code, but got: %v\n", rw.Code)
	}
}

func TestParamters(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/?name=test.fghg&lenght=56&breadth=90", nil)
    rw := httptest.NewRecorder()
    http.DefaultServeMux.ServeHTTP(rw, req)
    expected := "Invalid Extension"
    actual := strings.Contains(rw.Body.String(),"Invalid Extension")
    if actual == false {
        t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
    }
 }

func TestParams(t *testing.T){
    req := httptest.NewRequest(http.MethodGet, "/?",nil)
    rw := httptest.NewRecorder()
    http.DefaultServeMux.ServeHTTP(rw,req)
    expected := "Invalid name"
    actual := strings.Contains(rw.Body.String(),"Invalid name")
    if actual == false {
       t.Errorf("Error")
    }
    fmt.Println(expected)
}


