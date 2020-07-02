package main

import(
	"fmt"
	"crypto/sha1"
	"github.com/wuriyanto48/go-pbkdf2"
)

func main(){
	pass := p.NewPassword(sha1.New, 8, 32,15000)
	hashed := pass.HashPassword("123456")
	fmt.Println(hashed.CipherText)
	fmt.Println(hashed.Salt)
}