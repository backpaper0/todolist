package inmemory

import (
	"reflect"
	"testing"

	"github.com/backpaper0/todolist/domain"
)

func TestAddAndGetAll(t *testing.T) {
	todolist := NewTodolist()

	todolist.Add("あれをやる")
	todolist.Add("これをやる")

	actual := todolist.GetAll()
	expected := []domain.Todo{
		{Id: "1", Task: "あれをやる", Done: false},
		{Id: "2", Task: "これをやる", Done: false},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected is %v but actual is %v", expected, actual)
	}
}

func TestUpdate(t *testing.T) {
	todolist := NewTodolist()

	todolist.Add("あれをやる")
	todolist.Add("これをやる")

	todolist.Update("1", true)
	todolist.Update("2", true)

	actual := todolist.GetAll()
	expected := []domain.Todo{
		{Id: "1", Task: "あれをやる", Done: true},
		{Id: "2", Task: "これをやる", Done: true},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected is %v but actual is %v", expected, actual)
	}
}

func TestClearAllDone(t *testing.T) {
	todolist := NewTodolist()

	todolist.Add("あれをやる")
	todolist.Add("これをやる")
	todolist.Add("それをやる")
	todolist.Add("どれをやる")

	todolist.Update("2", true)
	todolist.Update("4", true)

	todolist.ClearAllDone()

	actual := todolist.GetAll()
	expected := []domain.Todo{
		{Id: "1", Task: "あれをやる", Done: false},
		{Id: "3", Task: "それをやる", Done: false},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected is %v but actual is %v", expected, actual)
	}
}
