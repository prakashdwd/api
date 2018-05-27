package main
import (
  
	"log"
	"net/http"
   
)





func showImage(w http.ResponseWriter, r *http.Request){
  
}   

func main() {
	
	http.HandleFunc("/", showImage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
