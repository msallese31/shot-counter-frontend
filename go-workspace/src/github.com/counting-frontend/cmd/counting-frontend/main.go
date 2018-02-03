package main

import "fmt"
import "github.com/counting-frontend/handler"
import "net/http"

func main() {
	fmt.Println("Starting server")
	// TODO: Chain handlers and have a logging handler
	http.Handle("/", &handler.RequestHandler{})
	http.ListenAndServe(":8080", nil)
}
