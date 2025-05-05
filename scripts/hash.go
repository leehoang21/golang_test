package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"

)

func main() {
	//hashed, _ := bcrypt.GenerateFromPassword([]byte("<password>"), bcrypt.DefaultCost)
	hashed, _ := bcrypt.GenerateFromPassword([]byte("Admin123@"), bcrypt.DefaultCost)
	fmt.Println(string(hashed))

}
