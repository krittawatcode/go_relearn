package hr

// struct (record)
type Person struct {
	Name string
	Age  int
}

/* anonymous field */
type Employee struct {
	Designation string
	Person
}
