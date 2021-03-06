package api

import (
	"encoding/json"
	"text/template"
	"net/http"
	"fmt"
)

type Task struct {
	Name string
	Done bool
}

func Init() {
	http.HandleFunc("/", Healthcheck)
	http.HandleFunc("/tasks", TaskHandler2)
	http.ListenAndServe(":8888", nil)
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	task := Task {
		Name: "Yup",
		Done: true,
	}
	
	t := template.Must(template.ParseFiles("templates/task.html"))
	t.Execute(w, task)
}

func TaskHandler2(w http.ResponseWriter, r *http.Request) {
	task := Task {
		Name: "Yup",
		Done: false,
	}

	j, _ := json.Marshal(task)
	w.Write(j)
}


func GetHttp(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic("Deu ruim na requisição!")
	}

	_showHttpResponse(res)
}

func UnsafeGetHttp(url string) {
	res, _ := http.Get(url)

	_showHttpResponse(res)
}

func _showHttpResponse(response *http.Response) {
	fmt.Println(response.Body)
}
