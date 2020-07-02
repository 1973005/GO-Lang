package main

import (
	"fmt"
	"github.com/apiang/gopackage-tutorial/config"
)

func main(){

	fmt.Println("Package Tutorial")
	
	pc := config.GetPostgresconnection()

	fmt.Println(pc)
}