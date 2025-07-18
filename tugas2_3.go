package main

import (
	"errors"
	"fmt"
)

// Fungsi untuk mencari bilangan genap dari slice
func cariGenap(data []int) ([]int, error) {
	var hasil []int
	for _, v := range data {
		if v%2 == 0 {
			hasil = append(hasil, v)
		}
	}
	if len(hasil) == 0 {
		return nil, errors.New("tidak ada bilangan genap yang ditemukan")
	}
	return hasil, nil
}

func main() {
	slice := []int{1, 3, 5, 7, 8, 10}
	genap, err := cariGenap(slice)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("bilangan genap:",
			genap)
	}
}
