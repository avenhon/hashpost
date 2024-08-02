package main

import "fmt"

func main() {
	var choice uint

	fmt.Println("Want to register or login?")
	fmt.Println("print 1 to register and 2 to login")
	fmt.Scan(&choice)

	switch choice {
		case(1): {
			fmt.Println("You select register")
		}
		case(2): {
			fmt.Println("You select login")
		}
	}
}