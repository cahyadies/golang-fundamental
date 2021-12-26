package main

import (
	"fmt"
	"golang-fundamental/calculation"
)

func main() {
	fmt.Println("Hello World")
	result := calculation.Add(22, 9)

	fmt.Println(result)
}

// Import digunakan untuk
// 1. Memanggil standar library
// 2. Memanggil package berbeda
// 3. Memanggil library third-party

// Package Main > Package Excetuable, harus mengandung Func Main
// Func Main, method yang akan dieksekusi pada package main
