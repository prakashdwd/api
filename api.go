package main
import ( 
  "log"
  "net/http"
  "strings"
  "image"
  "os"
)

var imageTemplate string = `<!DOCTYPE html><html lang="en"><head></head><body><img src="data:image/jpg;base64,{{.Image}}"></body>`
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
            return false
        }
    }
    return true
}


func showImage(w http.ResponseWriter, r *http.Request){
  name := r.URL.Query()["name"]
  length := r.URL.Query()["length"]
  breadth := r.URL.Query()["breadth"]

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
          return
        }
     }
     file_name_length := strings.Index(name[0],".") 
     file_name := string(name[0][:file_name_length])

     path := "images/" + file_name + "_" + length[0] + "_" + breadth[0]  + format
     original_path := "images/" + file_name + format

     if !Exists(path){
         if !Exists(original_path){
            ErrorHandler(w , "No such image")
            return
         }else{
            CreateImage(w , file_name , length[0] , breadth[0] , format)
         }
     }

     file , err := os.Open(path)
   if err != nil {
      ErrorHandler(w,"Cannot open requested image!")
      return
   }

   img,_ , err1 := image.Decode(file)
	if err1 != nil {
		ErrorHandler(w,"Error decoding image!")
		return
	}
   WriteImageWithTemplate(w, &img , format , imageTemplate)
}


func main() {
  http.HandleFunc("/", showImage)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
