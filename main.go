package main

import (
	"fmt"
	"pertama/calculation"
)

//untuk menjalankan func diluar package main/beda package menggunakan import
//"namamodule/package" seperti contoh diatas pertama/calculation
func main() {
	fmt.Println("Halo Golang!")

	sentence := Testing()

	fmt.Println(sentence)

	result := calculation.Add(2, 3)
	fmt.Println(result)
}
