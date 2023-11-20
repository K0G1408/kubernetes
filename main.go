package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es una aplicacion usando Kubernetes")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("Corriendo en puerto 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}