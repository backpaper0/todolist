package todolist

import "strconv"

type Todolist struct {
	value []*Todo
}

type Todo struct {
	Id   string
	Task string
	Done bool
}

func New() *Todolist {
	todolist := &Todolist{}
	todolist.value = make([]*Todo, 0)
	return todolist
}

func (todolist *Todolist) Add(task string) {
	id := strconv.Itoa(len(todolist.value) + 1)
	todo := &Todo{id, task, false}
	todolist.value = append(todolist.value, todo)
}

func (todolist *Todolist) GetAll() []Todo {
	list := make([]Todo, 0, len(todolist.value))
	for _, todo := range todolist.value {
		list = append(list, *todo)
	}
	return list
}

func (todolist *Todolist) Update(id string, done bool) {
	for _, todo := range todolist.value {
		if todo.Id == id {
			todo.Done = done
		}
	}
}

func (todolist *Todolist) ClearAllDone() {
	list := make([]*Todo, 0)
	for _, todo := range todolist.value {
		if !todo.Done {
			list = append(list, todo)
		}
	}
	todolist.value = list
}
