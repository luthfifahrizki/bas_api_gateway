package main

import (
	"api_gateway/usecase"
	"fmt"
)

func main() {
	login := usecase.NewLogin()
	auth := login.Autentikasi("admin", "admin123")
	fmt.Println(auth)
}
