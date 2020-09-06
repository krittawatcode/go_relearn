package main

import (
	"encoding/json"
	"fmt"
)

var data1 = `{
	"userId": 1,
	"id": 1,
	"title": "delectus aut autem",
	"completed": false
  }`

var data2 = `[{
	"userId": 1,
	"id": 1,
	"title": "delectus aut autem",
	"completed": false
	},
	{
	"userId": 1,
	"id": 2,
	"title": "quis ut nam facilis et officia qui",
	"completed": false
	}]`

type Todo struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func main() {

	/* ---- Single data ---- */
	data1Struct := Todo{}
	fmt.Println(data1Struct)
	v := &data1Struct
	json.Unmarshal([]byte(data1), v)
	fmt.Println(data1Struct)

	fmt.Println("--------------------------------------------")

	/* ---- Multiple data ---- */
	data2Struct := []Todo{}
	fmt.Println(data2Struct)
	v2 := &data2Struct
	json.Unmarshal([]byte(data2), v2)
	fmt.Println(data2Struct)
	fmt.Println(len(data2Struct))
}
