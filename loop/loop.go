package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	for key, value := range [5]int{1, 2, 3, 4, 5} {
		fmt.Println(key, value)
	}

	for {
		fmt.Println("hey")
		break
	}
}
