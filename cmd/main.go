package main

import (
	"fmt"

	"github.com/ernstlegaspi/todolist/cmd/api"
)

func main() {
	server := api.InitServer(":3001")

	if err := server.RunAPI(); err != nil {
		fmt.Println(err)
		fmt.Println("Error line 14 main.go")
		return
	}
}
