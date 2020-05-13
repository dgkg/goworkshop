package main

import "fmt"

// TODO: modify the signature methods functions on the User type so this use a coherent pointer.

type User struct {
	firstName, lastName string
}

func NewUser(fn, ln string) User {
	return User{
		firstName: fn,
		lastName:  ln,
	}
}

func (u User) SayHello() {
	fmt.Println("Hi! My name is " + u.firstName)
}

func (u *User) changeLastName(ln string) {
	u.lastName = ln
}

func main() {
	u := NewUser("Robert", "Pick")
	u.SayHello()
}
