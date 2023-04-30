package todolist

import "strconv"

type Todo struct {
	Id   string
	Task string
	Done bool
}

var todolist = make([]*Todo, 0)

func Add(task string) {
	id := strconv.Itoa(len(todolist) + 1)
	todo := &Todo{id, task, false}
	todolist = append(todolist, todo)
}

func GetAll() []*Todo {
	return todolist
}

func Update(id string, done bool) {
	for _, todo := range todolist {
		if todo.Id == id {
			todo.Done = done
		}
	}
}

func Delete(id string) {
	newTodolist := make([]*Todo, 0)
	for _, todo := range todolist {
		if todo.Id != id {
			newTodolist = append(newTodolist, todo)
		}
	}
	todolist = newTodolist
}
