package main

import "fmt"

func main() {
	fmt.Println(createUsers())
	fmt.Println(createCats())
	fmt.Println(createDogs())
}

// TODO: help me to fix me.
func createUsers() [3]string {
	return [3]string{
		"Chewbarka",
		"Jimmy Chew",
		"Bark Wahlberg",
		"Droolius Caesar",
		"Bark Twain",
	}
}

// TODO: help me to fix me.
func createCats() {
	var cats [2]string
	cats[1] = "Abbie"
	cats[2] = "Oward"
	return cats
}

// TODO: help me to fix me.
func createDogs() (res [2]string) {
	var dogs [2]string
	dogs[0] = 'Amber'
	dogs[1] = 'Stoot'
}
