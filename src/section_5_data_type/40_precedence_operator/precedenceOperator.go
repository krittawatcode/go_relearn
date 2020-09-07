package main

import "fmt"

func main() {
	/*
	   if you are not study in programatic field then this will help you walk through alot
	   https://www.tutorialspoint.com/go/go_operators.htm // Bitwise Operators part
	   https://code.tutsplus.com/articles/understanding-bitwise-operators--active-11301 //

	*/

	// %
	/*
		like in mathematic we use 'and(&)' to compare and if both have same value then we return 1 if not return 0
		ex A = true, B = false
		then A & B = false
		In byte type we will compare every single byte in same position
		and return true false value in that position with 1 mean true and 0 mean false
		ex 5 & 2 it mean -> 0000 0101
							0000 0010
					if we compare it last digit "1 and 0" then result is false then we put 0 there
					so if we compare every single digit result will be like below
							0000 0000 -> and if we convert base2 to base10 then value is 0
	*/
	fmt.Println("5 % 2 : ", 5%2)     // 1
	fmt.Println("5 % -2 :", 5%-2)    // 1
	fmt.Println("-5 % -2 : ", -5%-2) // -1

	fmt.Println("5 & 2 : ", 5&2)   // 0
	fmt.Println("5 & 5 : ", 5&5)   // 5
	fmt.Println("5 &^ 2 : ", 5&^2) // 5
}
