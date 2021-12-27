package main

import "fmt"

func main() {

	// for biasa
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i, "Halo Dunia")
	// }

	// for gaya while
	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Halo", i)
	// 	i++
	// }

	//for each
	title := "Golang the best language"

	// for index, letter := range title {
	// 	fmt.Println("Index :", index, "Letter :", string(letter))
	// }
	// _ -> variable penampung, jika variable tidak dipakai
	for _, letter := range title {
		fmt.Println("Letter :", string(letter))
	}
}
