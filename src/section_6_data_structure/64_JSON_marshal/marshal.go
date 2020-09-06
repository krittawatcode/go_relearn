package main

import (
	"encoding/json"
	"fmt"
)

var data = `[{
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
	},
	{
	"userId": 15,
	"id": 20,
	"completed": false
	},
	{
	"userId": 15,
	"id": 20
	}]`

// struct with json output format
// !!!! with omitempty -> title and bool
type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed *bool  `json:"completed,omitempty"`
}

func main() {

	dataStruct := []Todo{}
	v := &dataStruct
	json.Unmarshal([]byte(data), v)
	// dataStruct[0].Completed = true
	// dataStruct[1].Completed = true

	// dataStruct --> std.Output
	// func Marshal(v interface{}) ([]byte, error) {...}

	/* ---- Marshal ---- */

	// result, err := json.Marshal(dataStruct)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(string(result))

	/* ----  MarshalIndent ----
	func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {...}
	*/
	result, err := json.MarshalIndent(dataStruct, "", "    ") // for make readable json format
	if err != nil {
		return
	}

	fmt.Println(string(result))
}
