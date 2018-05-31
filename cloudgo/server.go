package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello! This is my first cloud web server"
	})
	m.Run()
}
