package main

import "fmt"

type User struct {
	FirstName, LastName string
}

// userList is the reprensentation of the user list.
// It has the uuid in the key and the value is the corresponding User.
var userList map[string]User

func main() {

	u := User{"Rob", "Pike"}

	// TODO: fixme !
	userList["DA92E99D-0C00-49E6-9343-FA89DA0670C6"] = u

	// TODO: implement the error check and log it.
	AddUser("DA92E99D-0C00-49E6-9343-FA89DA0670C6", u)

	// TODO: implement the error check and log it.
	DeleteUser("63FC8E21-240A-4E75-98A2-962A5ED67177")

	// this should display the user list
	fmt.Println(userList)
}

// AddUser is adding a new user in the userList.
// this send an error if the user allready exsits in the userList.
func AddUser(uuid string, u User) error {
	// TODO: implement this function
	// this should add a new user in the userList map
	// if the given uuid allready exists in the userList key this should send an error.
	// otherwise this should send a nil error value
}

// DeleteUser is deleting the given uuid user from the userList.
// if the user don't exists in the userList this function sends an error.
func DeleteUser(uuid string) error {
	// TODO: implement this function.
	// this should delete in the userList the given uuid.
	// if the given uuid is not found this should send an error
	// otherwise this should send a simple nil error value.
}
