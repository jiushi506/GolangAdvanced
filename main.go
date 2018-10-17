package main

import "fmt"

func main() {
	a := new(int64)
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(1<<7)
}
