package inmemory

import (
	"strconv"

	"github.com/backpaper0/todolist/domain"
)

type InMemoryTodolist struct {
	value []*domain.Todo
}

func NewTodolist() domain.Todolist {
	todolist := &InMemoryTodolist{}
	todolist.value = make([]*domain.Todo, 0)
	return todolist
}

func (todolist *InMemoryTodolist) Add(task string) {
	id := strconv.Itoa(len(todolist.value) + 1)
	todo := &domain.Todo{
		Id:   id,
		Task: task,
		Done: false,
	}
	todolist.value = append(todolist.value, todo)
}

func (todolist *InMemoryTodolist) GetAll() []domain.Todo {
	list := make([]domain.Todo, 0, len(todolist.value))
	for _, todo := range todolist.value {
		list = append(list, *todo)
	}
	return list
}

func (todolist *InMemoryTodolist) Update(id string, done bool) {
	for _, todo := range todolist.value {
		if todo.Id == id {
			todo.Done = done
		}
	}
}

func (todolist *InMemoryTodolist) ClearAllDone() {
	list := make([]*domain.Todo, 0)
	for _, todo := range todolist.value {
		if !todo.Done {
			list = append(list, todo)
		}
	}
	todolist.value = list
}
