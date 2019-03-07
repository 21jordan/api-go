package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
	fmt.Println("ejecutando servidor de go en le puerto 8080")

}
