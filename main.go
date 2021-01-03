package main

import (
	"fmt"
	"log"
	"net/http"

	"grandeurSeason/routes"
)

func main() {
	r := routes.NewRouter()
	fmt.Println("Start")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
