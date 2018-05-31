package main

import (
	"golang-mysql/service"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", service.NewServer())
}
