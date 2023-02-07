package main

import (
	"fmt"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("building MongoDB Api");
	fmt.Println("Server is getting Started...");

	r := router.Router();

	http.ListenAndServe(":8080", r)

	fmt.Println("Listening at port 8080")
}