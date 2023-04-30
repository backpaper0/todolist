package web

import (
	"html/template"
	"net/http"
	"strconv"

	_ "embed"

	"github.com/backpaper0/todolist"
)

type Web struct {
	repos *todolist.Todolist
}

var tmpl *template.Template

//go:embed index.html
var htmlSource string

func init() {
	tmpl = template.Must(template.New("index").Parse(htmlSource))
}

func NewWeb() *Web {
	return &Web{
		repos: todolist.New(),
	}
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
		//TODO エラー
		rw.WriteHeader(400)
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
		//TODO エラー
		rw.WriteHeader(400)
		return
	}
	id := ids[0]
	done, err := strconv.ParseBool(dones[0])
	if err != nil {
		rw.WriteHeader(400)
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
