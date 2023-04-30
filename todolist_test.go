package todolist

import (
	"reflect"
	"testing"
)

func TestAddAndGetAll(t *testing.T) {
	todolist := New()

	todolist.Add("あれをやる")
	todolist.Add("これをやる")

	actual := todolist.GetAll()
	expected := []Todo{
		{"1", "あれをやる", false},
		{"2", "これをやる", false},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected is %v but actual is %v", expected, actual)
	}
}

func TestUpdate(t *testing.T) {
	todolist := New()

	todolist.Add("あれをやる")
	todolist.Add("これをやる")

	todolist.Update("1", true)
	todolist.Update("2", true)

	actual := todolist.GetAll()
	expected := []Todo{
		{"1", "あれをやる", true},
		{"2", "これをやる", true},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected is %v but actual is %v", expected, actual)
	}
}

func TestDelete(t *testing.T) {
	todolist := New()

	todolist.Add("あれをやる")
	todolist.Add("これをやる")

	todolist.Delete("2")

	actual := todolist.GetAll()
	expected := []Todo{
		{"1", "あれをやる", false},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected is %v but actual is %v", expected, actual)
	}
}
