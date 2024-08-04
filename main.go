package main

import (
	"fmt"
)

type userAccount struct {
	nickname string
	password string
	email string
	age uint
}

func main() {
	var choice uint
	userProfiles := make(map[string]userAccount)

	var currentAccount userAccount

	for {
		for currentAccount == (userAccount{}) {
			fmt.Println("Want to register or login?")
			fmt.Printf("print 1 to register and 2 to login: ")
			fmt.Scan(&choice)
		
			switch choice {
				case(1): {
					userRegister(&userProfiles, &currentAccount)
				}
				case(2): {
					fmt.Println("You select login")
				}
			}
		}
	
	
		switch showMenu() {
			case(3): {
				currentAccount = userAccount{}
			}
			case(4): {
				return
			}
		}
	}
}

func userRegister(userProfiles *map[string]userAccount, currentAccount *userAccount) {
	var newUserAccount userAccount
	var nickname string
	var password string
	var email string
	var age uint

	var choice int

	for {
		fmt.Println("You selected register")
		fmt.Printf("Print your nickname: ")
		fmt.Scan(&nickname)
	
		fmt.Printf("Print your password: ")
		fmt.Scan(&password)
	
		fmt.Printf("Print your email: ")
		fmt.Scan(&email)
	
		fmt.Printf("Print your age: ")
		fmt.Scan(&age)

		fmt.Printf("%v your password is %v, email: %v and age %v, confirm registration?\n", nickname, password, email, age)
		fmt.Printf("To confirm the action, print 1. Else print 2: ")
		fmt.Scan(&choice)
		
		if choice == 1 {
			break
		} else if choice == 2 {
			continue
		}
	}


	newUserAccount.nickname = nickname
	newUserAccount.password = password
	newUserAccount.email = email
	newUserAccount.age = age

	*currentAccount = newUserAccount
	(*userProfiles)[nickname] = newUserAccount
}

func showMenu() uint {
	var choice uint
	fmt.Println("Menu of actions:")
	fmt.Printf("To exit from account print 3\n")
	fmt.Printf("To exit from program print 4\n")
	fmt.Scan(&choice)
	return choice
}