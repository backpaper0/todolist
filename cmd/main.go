package main

import (
	"net/http"

	"github.com/backpaper0/todolist/domain"
	"github.com/backpaper0/todolist/web"
)

func main() {
	repos := domain.NewTodolist()
	w := web.NewWeb(repos)
	server := &http.Server{
		Handler: w.Handler,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
