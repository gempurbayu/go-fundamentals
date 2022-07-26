package main

import (
	"fmt"
	"pertama/calculation"
	"pertama/conditional"
)

//untuk menjalankan func diluar package main/beda package menggunakan import
//"namamodule/package" seperti contoh diatas pertama/calculation
func main() {
	fmt.Println("Halo Golang!")

	//sentence := Testing()

	//fmt.Println(sentence)

	result := calculation.Add(2, 3)
	fmt.Println(result)

	//example definiton variable
	var name string = "Gempur"
	var kota string

	//penggunaan := untuk definisi variable yang sudah ada valuenya (tanpa harus ada var dan nama tipedata)
	// age:= 20 => var age int = 20
	age := 20
	//tidak harus dengan := karena sudah didefinisikan diatas
	age = 30

	kota = "Bandung"
	fmt.Println(name, kota, age)

	usia := conditional.TestIf()
	huruf := conditional.SwitchCase()

	fmt.Println(usia, huruf)

	conditional.ForEx()

}
