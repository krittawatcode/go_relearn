package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	resp, err := http.Get("http://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}
	todoDecoder := json.NewDecoder(resp.Body)
	todos := []Todo{}
	todoDecoder.Decode(&todos)
	resp.Body.Close()
	// todoEncoder := json.NewEncoder(os.Stdout)
	// todoEncoder.Encode(todos)

	indexTemplateBytes, err := ioutil.ReadFile("D:/Learn/GO/src/section_6_data_structure/69_http_template/index.html")
	if err != nil {
		return
	}

	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateBytes)))
	indexTemplate.Execute(os.Stdout, todos)
	// when finish run to create new html
	// go run .\src\section_6_data_structure\69_http_template\http_template.go > todo.html

}
