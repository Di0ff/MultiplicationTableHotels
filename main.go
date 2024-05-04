package main

import (
	"fmt"
)

func main() {
	MultiplicationTable(5)
}

func MultiplicationTable(n int) {
	fmt.Print("  ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%3d", i)
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%2d", i)
		for j := 1; j <= n; j++ {
			fmt.Printf("%3d", i*j)
		}
		fmt.Println()
	}
}
