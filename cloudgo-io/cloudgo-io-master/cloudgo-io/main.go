package main

import (
	"cloudgo-io/router"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", router.NewServer())
}
