package main

import "fmt"

func main() {
	//fn := anonymous()
	//fn("hello there")
}

func func1() {
	fmt.Println("this is func1")
}

func func2(i int)  {
	fmt.Println("func2: ", i)
}

func func3() int {
	return 10
}

func func4(i int) int {
	return i
}

func func5() (string, string) {
	return "vikram", "jakhar"
}

type MyType string

// pointer
func (t *MyType) func6(name string) (string, string) {
	fmt.Println("MyType: ", t)
	fmt.Println(name)
	return "a", "b"
}

func anonymous() (func(string)){
	return func(s string) {
		fmt.Println(s)
	}
}