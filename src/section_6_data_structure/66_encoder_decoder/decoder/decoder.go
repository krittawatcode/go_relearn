package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type Users []struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

func main() {

	resp, err := http.Get("http://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}

	// json.decode (json --> dataStruct)
	jsonDecoder := json.NewDecoder(resp.Body)
	dataStruct := Users{}
	jsonDecoder.Decode(&dataStruct)

	dataStruct[0].Name = "Krittawat"

	result, err := json.MarshalIndent(dataStruct, "", "    ") // for make readable json format
	if err != nil {
		return
	}

	fmt.Println(string(result))
}
