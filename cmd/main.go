package main

import (
	"net/http"

	"github.com/backpaper0/todolist/web"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", web.GetAll)
	handler.HandleFunc("/add", web.Add)
	handler.HandleFunc("/update", web.Update)
	handler.HandleFunc("/clearAllDone", web.ClearAllDone)

	server := &http.Server{
		Handler: handler,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
