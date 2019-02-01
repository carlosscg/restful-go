package main

import (
	"fmt"
	"net/http"
	"os"
	"restful-api/first/handlers"
)

func main() {
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	error := http.ListenAndServe("localhost:1111", nil)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

}
