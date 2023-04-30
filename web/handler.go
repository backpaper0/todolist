package web

import (
	"html/template"
	"net/http"
	"strconv"

	_ "embed"

	"github.com/backpaper0/todolist"
)

type Web struct {
	repos   *todolist.Todolist
	Handler *http.ServeMux
}

var tmpl *template.Template

//go:embed index.html
var htmlSource string

func init() {
	tmpl = template.Must(template.New("index").Parse(htmlSource))
}

func NewWeb() *Web {
	web := &Web{
		repos: todolist.New(),
	}
	handler := http.NewServeMux()
	handler.HandleFunc("/", web.GetAll)
	handler.HandleFunc("/add", web.Add)
	handler.HandleFunc("/update", web.Update)
	handler.HandleFunc("/clearAllDone", web.ClearAllDone)
	web.Handler = handler
	return web
}

func (web *Web) GetAll(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		rw.WriteHeader(405)
		return
	}
	data := make(map[string]interface{})
	data["List"] = web.repos.GetAll()
	tmpl.Execute(rw, data)
}

func (web *Web) Add(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(405)
		return
	}
	req.ParseForm()
	task := req.Form["task"]
	if len(task) == 0 {
		data := make(map[string]interface{})
		data["ErrorMessage"] = "タスクを入力してください。"
		data["List"] = web.repos.GetAll()
		tmpl.Execute(rw, data)
		return
	}
	web.repos.Add(task[0])
	rw.Header().Add("Location", "/")
	rw.WriteHeader(303)
}

func (web *Web) Update(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(405)
		return
	}
	req.ParseForm()
	ids := req.Form["id"]
	dones := req.Form["done"]
	if len(ids) == 0 || len(dones) == 0 {
		data := make(map[string]interface{})
		data["ErrorMessage"] = "不正な入力値です。"
		data["List"] = web.repos.GetAll()
		tmpl.Execute(rw, data)
		return
	}
	id := ids[0]
	done, err := strconv.ParseBool(dones[0])
	if err != nil {
		data := make(map[string]interface{})
		data["ErrorMessage"] = "不正な入力値です。"
		data["List"] = web.repos.GetAll()
		tmpl.Execute(rw, data)
		return
	}
	web.repos.Update(id, done)
	rw.Header().Add("Location", "/")
	rw.WriteHeader(303)
}

func (web *Web) ClearAllDone(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(405)
		return
	}
	req.ParseForm()
	web.repos.ClearAllDone()
	rw.Header().Add("Location", "/")
	rw.WriteHeader(303)
}
