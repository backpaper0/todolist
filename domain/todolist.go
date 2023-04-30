package domain

type Todolist interface {
	Add(task string)
	GetAll() []Todo
	Update(id string, done bool)
	ClearAllDone()
}

type Todo struct {
	Id   string
	Task string
	Done bool
}
