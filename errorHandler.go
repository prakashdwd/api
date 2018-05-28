package main
import (   	
	"log"
	"net/http"
	"html/template"
)

var ErrorTemplate string = `<!DOCTYPE html><html lang="en"><head><{{.Message}}></head><body></body>`

func ErrorHandler (w http.ResponseWriter , message string) {
    if new_template, err := template.New("text").Parse(ErrorTemplate); err != nil {
        log.Println("unable to parse image template.")
    } else {
        data := map[string]interface{}{"Message": message}
        if err = new_template.Execute(w, data); err != nil {
            log.Println("unable to execute template.")
        }
    }
}