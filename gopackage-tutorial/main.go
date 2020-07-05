package main

import (
	"fmt"

	"go-work/src/apiang/gopackage-tutorial/config"
)

func main() {

	fmt.Println("Package Tutorial")

	pc := config.GetPostgresconnection()

	fmt.Println(pc)
}
