package main

import "fmt"

type Sword struct {
	name string
}

func (s Sword) detail() string {
	return "Edge cut"
}

func (s Sword) cost() int {
	return 1000000000
}

type Bow struct {
	name string
}

func (s Bow) detail() string {
	return "Piercing Arrow"
}

func (s Bow) cost() int {
	return 9999999999
}

type Item interface {
	cost() int
}

type Weapon interface {
	Item // use other interface -> satisfy interface
	detail() string
}

func attack(w Weapon) {
	fmt.Printf("attack by : %s. Item cost : %d\n", w.detail(), w.cost())
}

/* attack by sword and bow */
func main() {
	sword := Sword{name: "Muramasa-Sword"}
	bow := Bow{name: "Gandiva-Bow"}

	var w Weapon
	w = sword
	attack(w)

	attack(sword)
	attack(bow)
}
