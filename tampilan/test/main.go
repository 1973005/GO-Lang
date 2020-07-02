package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// get an image file
	res, err := http.Get("https://bit.ly/2IRnmVm")

	// check for response error
	if err != nil {
		log.Fatal(err)
	}

	// read response data
	data, _ := ioutil.ReadAll(res.Body)

	// close response body
	res.Body.Close()

	// check if `Content-Type` is `image/jpeg`
	if res.Header.Get("Content-Type") == "image/jpeg" {
		ioutil.WriteFile("keyboard.jpeg", data, 0777)
		fmt.Println("File saved!")
	} else {
		log.Fatal("Error: Response is not an image.")
	}
}
