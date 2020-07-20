package main

import (
	"fmt"
	"gopackage-tutorial/config"
)

func main() {

	fmt.Println("Package Tutorial")

	pc := config.GetPostgresconnection()

	fmt.Println(pc)
}
