package main

import (
	"net/http"

	"github.com/backpaper0/todolist/web"
)

func main() {
	w := web.NewWeb()

	handler := http.NewServeMux()
	handler.HandleFunc("/", w.GetAll)
	handler.HandleFunc("/add", w.Add)
	handler.HandleFunc("/update", w.Update)
	handler.HandleFunc("/clearAllDone", w.ClearAllDone)

	server := &http.Server{
		Handler: handler,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
