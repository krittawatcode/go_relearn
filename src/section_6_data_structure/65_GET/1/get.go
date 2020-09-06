package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {

	resp, err := http.Get("http://jsonplaceholder.typicode.com/todos")
	// fmt.Println(err)
	// fmt.Println(resp)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	dataStruct := []Todo{}
	v := &dataStruct
	json.Unmarshal(body, v)

	// fmt.Println(dataStruct)

	result, err := json.MarshalIndent(dataStruct, "", "    ")
	if err != nil {
		return
	}

	fmt.Println(string(result))
}
