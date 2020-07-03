package main

import (
	"fmt"
	"os"
)

func main() {

	postgres := os.Getenv("POSTGRES")
	mysql := os.Getenv("MYSQL")

	fmt.Println("postgres Config : ", postgres)
	fmt.Println("MySQL Config : ", mysql)
}
