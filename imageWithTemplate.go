package main
import (
    "log"
    "net/http"
    "html/template"
    "bytes"
    "encoding/base64"
    "image/jpeg"
    "image/png"
    "image"
)
func WriteImageWithTemplate(w http.ResponseWriter, img *image.Image , format , imageTemplate string) {
    buffer := new(bytes.Buffer)  
    switch format{
	case ".png":
		if err := png.Encode(buffer, *img); err != nil {
        	log.Println("unable to encode image.")
        }
	case ".jpeg":
		if err := jpeg.Encode(buffer, *img, nil); err != nil {
        	log.Println("unable to encode image.")
        }
	}
    encoded_string := base64.StdEncoding.EncodeToString(buffer.Bytes())
    if tmpl, err := template.New("image").Parse(imageTemplate); err != nil {
        log.Println("unable to parse image template.")
    } else {
        data := map[string]interface{}{"Image": encoded_string}
        if err = tmpl.Execute(w, data); err != nil {
            log.Println("unable to execute template.")
        }
    }
}

