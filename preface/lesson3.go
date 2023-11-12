package main

import "fmt"

func main() {
	var number int32
	var name string
	fmt.Println("Please enter a int number:")
	_, err := fmt.Scanln(&number)
	if err != nil {
		return
	}
	fmt.Println("Please enter a name:")
	_, err = fmt.Scanln(&name)
	if err != nil {
		return
	}
	fmt.Printf("number is: %d, name is %v", number, name)
}
