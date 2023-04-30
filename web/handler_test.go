package web

import (
	"io"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	w := NewWeb()
	server := httptest.NewServer(w.Handler)
	defer server.Close()

	client := server.Client()

	resp, _ := client.Get(server.URL)

	bs, _ := io.ReadAll(resp.Body)
	s := string(bs)
	if !strings.Contains(s, "TODOはありません。") {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	w := NewWeb()
	server := httptest.NewServer(w.Handler)
	defer server.Close()

	client := server.Client()

	form := url.Values{}
	form.Add("task", "タスク1")
	resp, _ := client.PostForm(server.URL+"/add", form)

	bs, _ := io.ReadAll(resp.Body)
	s := string(bs)
	if !strings.Contains(s, "タスク1") {
		t.Error(s)
	}
}

func TestUpdate(t *testing.T) {
	w := NewWeb()
	server := httptest.NewServer(w.Handler)
	defer server.Close()

	client := server.Client()

	form := url.Values{}
	form.Add("task", "タスク1")
	resp, _ := client.PostForm(server.URL+"/add", form)

	bs, _ := io.ReadAll(resp.Body)
	s := string(bs)
	if !strings.Contains(s, "タスク1") || !strings.Contains(s, "完了") {
		t.Error(s)
	}

	form = url.Values{}
	form.Add("id", "1")
	form.Add("done", "true")
	resp, _ = client.PostForm(server.URL+"/update", form)
	bs, _ = io.ReadAll(resp.Body)
	s = string(bs)
	if !strings.Contains(s, "タスク1") || !strings.Contains(s, "戻す") {
		t.Error(s)
	}

	form = url.Values{}
	form.Add("id", "1")
	form.Add("done", "false")
	resp, _ = client.PostForm(server.URL+"/update", form)
	bs, _ = io.ReadAll(resp.Body)
	s = string(bs)
	if !strings.Contains(s, "タスク1") || !strings.Contains(s, "完了") {
		t.Error(s)
	}
}

func TestClearAllDone(t *testing.T) {
	w := NewWeb()
	server := httptest.NewServer(w.Handler)
	defer server.Close()

	client := server.Client()

	tasks := []string{
		"タスク1",
		"タスク2",
		"タスク3",
		"タスク4",
	}
	for _, task := range tasks {
		form := url.Values{}
		form.Add("task", task)
		client.PostForm(server.URL+"/add", form)
	}

	ids := []string{
		"1",
		"3",
	}
	for _, id := range ids {
		form := url.Values{}
		form.Add("id", id)
		form.Add("done", "true")
		client.PostForm(server.URL+"/update", form)
	}

	form := url.Values{}
	resp, _ := client.PostForm(server.URL+"/clearAllDone", form)

	bs, _ := io.ReadAll(resp.Body)
	s := string(bs)
	if strings.Contains(s, "タスク1") ||
		!strings.Contains(s, "タスク2") ||
		strings.Contains(s, "タスク3") ||
		!strings.Contains(s, "タスク4") {
		t.Error(s)
	}
}
