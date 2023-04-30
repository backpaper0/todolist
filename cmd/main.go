package main

import (
	"net/http"

	"github.com/backpaper0/todolist/web"
)

func main() {
	w := web.NewWeb()
	server := &http.Server{
		Handler: w.Handler,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
