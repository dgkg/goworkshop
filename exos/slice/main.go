package main


func main() {
	
}

// User represent a user with a fristname and a lastname.
type User struct{
	FirstName, LastName string
}

// NewUser returns a new user initialize with the params.
// TODO: refactor the function signature with the right params in it (firstname and lastname)
// TODO: add a return value of the function in the function signature
// TODO: return a pointer of the initialized user 
func NewUser(){
	return // ???
}

// TODO: change this current list so we can add new users in it.
var userList [5]User{
	{"Lois","Brooks"},
	{"Gladys","Stevens"},
	{"Willie","Little"},
	{"Andre","Grant"},
	{"Constance","Graham"},
}


// GetUser is returning the current userList
// TODO: refactor the function signature.
func GetUser()[5]User{
	return userList
}

// AddUser is adding a new user in the userList.
// TODO: implement the function so the userList has an appended new User.
func AddUser(User){

}
