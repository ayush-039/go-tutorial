package main

import (
	"fmt"

	"github.com/ayush-039/go-tutorial/auth"
	"github.com/ayush-039/go-tutorial/user"
)

func main() {
	auth.LoginWithCredientials("ayush Barot", "ayush1234")
	session := auth.GetSession()
	fmt.Println("session:", session)

	user := user.User{
		Email: "ayushabarot3902@gmail.com",
		Name:  "Ayush Barot",
	}

	fmt.Println(user)
}
