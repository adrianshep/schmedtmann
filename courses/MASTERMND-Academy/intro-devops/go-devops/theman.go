package main

import "fmt"

// func imTheMan() string {
// 	return "I'm the man"
// }

func imTheMan(name string) string {
	return name + " is the man!"
}

func main() {
	fmt.Println(imTheMan("Ado"))
}
