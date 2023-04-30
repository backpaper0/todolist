package todolist

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	todolist = make([]*Todo, 0)

	Add("あれをやる")

	todo := todolist[len(todolist)-1]
	if todo.Task != "あれをやる" || todo.Done != false {
		t.Fail()
	}
}

func TestGetAll(t *testing.T) {
	todolist = []*Todo{
		{"1", "あれをやる", false},
		{"2", "これをやる", true},
	}

	actual := GetAll()
	expected := []*Todo{
		{"1", "あれをやる", false},
		{"2", "これをやる", true},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fail()
	}
}

func TestUpdate(t *testing.T) {
	todolist = []*Todo{
		{"1", "あれをやる", false},
		{"2", "これをやる", true},
	}

	Update("1", true)
	Update("2", false)

	expected := []*Todo{
		{"1", "あれをやる", true},
		{"2", "これをやる", false},
	}
	if !reflect.DeepEqual(todolist, expected) {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	todolist = []*Todo{
		{"1", "あれをやる", false},
		{"2", "これをやる", true},
	}

	Delete("2")

	expected := []*Todo{
		{"1", "あれをやる", false},
	}
	if !reflect.DeepEqual(todolist, expected) {
		t.Fail()
	}
}
