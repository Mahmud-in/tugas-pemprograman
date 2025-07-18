package main

import "fmt"

func swap(a, b *string) {
	*a, *b = *b, *a
}

func main() {
	str1 := "Hello"
	str2 := "world"
	swap(&str1, &str2)
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
}
