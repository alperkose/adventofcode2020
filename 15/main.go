package main

import "fmt"

func main() {
	fmt.Println(NewGame([]int{2, 1, 10, 11, 0, 6}).Guess(2020))
	fmt.Println(NewGame([]int{2, 1, 10, 11, 0, 6}).Guess(30000000))
}
