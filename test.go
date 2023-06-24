package main

import "fmt"

func main() {
	c := "兵"
	fmt.Println("abcde")
	fmt.Printf("%2s\n", c)
	fmt.Printf("%3s\n", c)

	fmt.Println("|------|")
	fmt.Println("|" + c + "  " + c + "|")
	fmt.Printf("|%6s|", " ")
}
