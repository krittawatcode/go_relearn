package main

import "fmt"

func main() {
	// map use hash table

	// ----- ex1 int map -----
	/*
		items := map[string]int{
			"pen":    4,
			"pencil": 5,
		}
	*/

	// ----- ex2 make int map
	/*
		orders := make(map[string]int)
		orders["pen"] = 16
		orders["pencil"] = 15
	*/

	// ----- delete map's entity -----
	/*
		items := map[string]int{
			"pen":    4,
			"pencil": 5,
		}
		delete(items, "pen")
	*/

	items := map[string]int{
		"pen":    4,
		"pencil": 5,
	}
	// fmt.Println(items)
	// fmt.Println(items["pen"])

	orders := make(map[string]int)
	orders["pen"] = 16
	orders["pencil"] = 15
	// fmt.Println(orders["pen"])

	// maps - use concept of pointer
	/*
		x := items
		x["ruler"] = 7
		fmt.Println("items : ", items)
		fmt.Println("x : ", x)
	*/

	x := items
	x["ruler"] = 7
	x = orders
	fmt.Println("items : ", items)
	fmt.Println("x : ", x)
	fmt.Println("items[\"pencil\"] : ", items["pencil"])
	fmt.Println("orders : ", orders)

	delete(items, "pencil")
	fmt.Println("items after delete : ", items)

	// check for have of have not ?
	v, ok := items["xxx"] // value and status
	if !ok {
		fmt.Println("not exist")
	} else {
		fmt.Println("Exist! :", v)
	}
	fmt.Println(items["xxx"])

	fmt.Println("-----------------------------")
	// map with range to prove map is a order data or not -->> not!!!
	itemss := map[string]int{
		"pen":    4,
		"pencil": 5,
		"a":      0,
		"b":      0,
		"c":      0,
	}

	// map don't be in order
	for k, v := range itemss {
		fmt.Println(k, v)
	}

	fmt.Println(len(itemss))
	// a := &itemss["pen"] -> cannot take the address of itemss["pen"]
}
