package main

import (
	"net/http"

	"github.com/backpaper0/todolist/domain/inmemory"
	"github.com/backpaper0/todolist/web"
)

func main() {
	repos := inmemory.NewTodolist()
	w := web.NewWeb(repos)
	server := &http.Server{
		Handler: w.Handler,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
