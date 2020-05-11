package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D"}
	b := make([]string, len(a))
	copy(b, a)
	a[0] = "Z"
	fmt.Println(a)
	fmt.Println(b)
}
