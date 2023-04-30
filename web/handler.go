package web

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/backpaper0/todolist"
)

var repos *todolist.Todolist
var tmpl *template.Template

func init() {
	repos = todolist.New()

	tmpl = template.Must(template.New("index").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>TODOリスト</title>
	</head>
	<body>
	<form method="POST" action="/add">
		<p>
			<input type="text" name="task" autofocus>
			<button>登録する</button>
		</p>
	</form>
	{{if .List}}
		<ul>
	{{end}}
	{{range .List}}
		<li>
			<form method="POST" action="/update">
				<button>{{if .Done}}戻す{{else}}完了{{end}}</button>
				<input type="hidden" name="id" value="{{.Id}}">
				<input type="hidden" name="done" value="{{not .Done}}">
				{{if .Done}}<del>{{end}}
				{{.Task}}
				{{if .Done}}</del>{{end}}
			</form>
		</li>
	{{else}}
		<p>TODOはありません。</p>
	{{end}}
	{{if .List}}
		</ul>
		<form method="POST" action="/clearAllDone">
			<p>
				<button>完了タスクをクリア</button>
			</p>
		</form>
		{{end}}
	</body>
	</html>
	`))
}

func GetAll(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		rw.WriteHeader(405)
		return
	}
	data := make(map[string]interface{})
	data["List"] = repos.GetAll()
	tmpl.Execute(rw, data)
}

func Add(rw http.ResponseWriter, req *http.Request) {
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
	repos.Add(task[0])
	rw.Header().Add("Location", "/")
	rw.WriteHeader(303)
}

func Update(rw http.ResponseWriter, req *http.Request) {
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
	repos.Update(id, done)
	rw.Header().Add("Location", "/")
	rw.WriteHeader(303)
}

func ClearAllDone(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(405)
		return
	}
	req.ParseForm()
	repos.ClearAllDone()
	rw.Header().Add("Location", "/")
	rw.WriteHeader(303)
}
