package main

import (
	"fmt"
)

func cariTerbesar(data []int) int {
	if len(data) == 0 {
		panic("Slice kosong, idak bisa mencari nilai terbesar")
	}
	maks := data[0]
	for _, v := range data {
		if v > maks {
			maks = v
		}
	}
	return maks
}

func main() {
	slice := []int{3, 7, 2, 10, 5}
	terbesar := cariTerbesar(slice)
	fmt.Println("Bilangan terbesar adalah:", terbesar)
}
