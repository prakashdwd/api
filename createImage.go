package main
import (
    "gopkg.in/gographics/imagick.v2/imagick"
    "strconv"
    "net/http"
)

func CreateImage(w http.ResponseWriter , name , length , breadth , format string) {
   
   imagick.Initialize()
   image_length, err1 := strconv.Atoi(length)
   image_breadth, err2 := strconv.Atoi(breadth)
   
   if err1 != nil {
      ErrorHandler(w , "Fatal error in atoi")
      return
   }
   if err2 != nil {
      ErrorHandler(w , "Fatal error in atoi")
      return
   }
   
   magic_wand := imagick.NewMagickWand()
   path := name + format
   err := magic_wand.ReadImage("images/" + path)
	if err != nil {
		ErrorHandler(w , "Error with magicWand")
    return
	}
	err = magic_wand.ResizeImage(uint(image_breadth), uint(image_length), imagick.FILTER_LANCZOS, 1)
	if err != nil {
		ErrorHandler(w , "Error with magicWand")
    return
	}
   magic_wand.WriteImage("images/" + name + "_" +  length + "_"  + breadth + format)
   magic_wand.Destroy()
   imagick.Terminate()
}

