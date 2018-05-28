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
 
    req = httptest.NewRequest(http.MethodGet, "/?",nil)
    http.DefaultServeMux.ServeHTTP(rw,req)
    expected = "Invalid name"
    actual = strings.Contains(rw.Body.String(),"Invalid name")
    if actual == false {
       t.Errorf("Error! Expected : %v , got : %v",actual,expected)
    }
    
    req = httptest.NewRequest(http.MethodGet , "/?name=tes.png&length=89&breadth=100",nil)
    http.DefaultServeMux.ServeHTTP(rw,req)
    expected = "No such image"
    actual = strings.Contains(rw.Body.String(),"No such image")
    if actual == false{
      t.Errorf("Error! Expected : %v , got : %v", actual , expected)
    }
    fmt.Println(rw.Body)
}

func TestCreateImage(t *testing.T){
   req := httptest.NewRequest(http.MethodGet , "/?name=test.jpeg&length=100&breadth=100",nil)
   rw := httptest.NewRecorder()
   http.DefaultServeMux.ServeHTTP(rw,req)
   expected := true
   actual := Exists("images/test_100_100.jpeg")
   if actual == false{
      t.Errorf("Error! Expected : %v , got : %v" , expected , actual)
   }
}




