package main

import (
	"flag"
	"fmt"
)

func main() {

	// config := flag.String("C", "1234", "your config")

	// flag.Parse()

	// myConfig := *config

	// fmt.Println(myConfig)

	username := flag.String("U", "", "Your Username")
	password := flag.String("P", "", "Your Password")

	flag.Parse()

	result := login(*username, *password)

	if result {
		fmt.Println("Login Success")
	} else {
		fmt.Println("Login Fail")
	}
}

func login(username, password string) bool {
	if username == "apiang" && password == "123456" {
		return true
	}
	return false
}

//go build
//cara runnya flag.exe
//-h untuk tau toolsnya
// contoh run (flag.exe -U apiang -P 123456)
