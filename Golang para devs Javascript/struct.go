package main

import "fmt"

type User struct {
	id   string
	name string
	age  int
}

//Definicao de funcoes
type InterfaceUser interface {
	PrintAge()
}

func (u User) PrintAge() {
	println(u.age)
}

func NewUser(u User) InterfaceUser {
	return u
}

func main() {

	user1 := NewUser(User{
		id:   "xdj3933-d9d9",
		name: "Maria",
		age:  21,
	})

	user2 := User{
		id:   "xdj3933-d9d9",
		name: "Joao",
		age:  20,
	}

	users := []InterfaceUser{
		user1,
		user2,
	}

	println("Users")

	fmt.Println(users)

	user1.PrintAge()
}
