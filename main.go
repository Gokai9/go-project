package main

import (
	"fmt"
	"gosite"
	"net/http"
)

func main() {
	webEngine := gosite.NewRouter()
	webEngine.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", req.URL.Path)
	})
	webEngine.Run(":8000")
}
