package main
import ( 
  "log"
  "net/http"
  "strings"
  "fmt"
)

var imageTemplate string = `<!DOCTYPE html><html lang="en"><head></head><body><img src="data:image/jpg;base64,{{.Image}}"></body>`




func showImage(w http.ResponseWriter, r *http.Request){
  name := r.URL.Query()["name"]
 // length := r.URL.Query()["length"]
 // breadth := r.URL.Query()["breadth"]

   var format string
   if name == nil {
      ErrorHandler(w , "Invalid name")
      return
   }

   if strings.Contains(name[0], "jpeg"){
      format = ".jpeg"
   }else{
      if strings.Contains(name[0], "png"){
        format = ".png"
        }else{
          ErrorHandler(w,"Invalid Extension")
          fmt.Println("herer-31")
          return
        }
      fmt.Println(format);
     }  
}


func main() {
  http.HandleFunc("/", showImage)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
