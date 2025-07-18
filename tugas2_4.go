package main

import (
	"errors"
	"fmt"
)

// Fungsi untuk mrncari bilangan ganjil dari slice
func cariGanjil(data []int) ([]int, error) {
	var hasil []int
	for _, v := range data {
		if v%2 != 0 {
			hasil = append(hasil, v)
		}
	}
	if len(hasil) == 0 {
		return nil, errors.New("tidak ada bilangan ganjil yang ditemukan")
	}
	return hasil, nil
}

func main() {
	slice := []int{2, 4, 6, 7, 9}
	ganjil, err := cariGanjil(slice)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Bilangan ganjil:", ganjil)
	}
}
