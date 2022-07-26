package data

import "fmt"

func Array() {

	// 3 adalah jumlah element, int tipe data element/array
	//var languages [3]string

	//definition 1
	//languages[0] = "GO"
	//languages[1] = "Ruby"
	//languages[2] = "Javascript"

	//definition 2
	//gunakan [...] jika tidak mau menentukan jumlah element
	languages := [...]string{"GO", "Ruby", "Javascript"}

	//fmt.Println(languages)

	for index, lang := range languages {
		fmt.Println("Index : ", index, " Language : ", lang)
	}

	//mengetahui panjang array
	length := len(languages)
	fmt.Println(length)
}

func Slice() {

	var Consoles []string

	Consoles = append(Consoles, "PS4")
	Consoles = append(Consoles, "Xbox")

	for _, cons := range Consoles {
		fmt.Println(cons)
	}
}

func Mapping() {

	//deklarasi variabel
	var myMap map[string]int

	//inisiasi value
	myMap = map[string]int{}

	//example 2
	myMap2 := map[string]string{"ruby": "is the best", "go": "the way"}

	myMap["ruby"] = 9
	myMap["go"] = 10

	fmt.Println(myMap2)
	fmt.Println(myMap["go"])

	for key, value := range myMap {
		fmt.Println("Key : ", key, " value : ", value)
	}

	fmt.Println("=======")

	delete(myMap, "ruby")
	for key, value := range myMap {
		fmt.Println("Key : ", key, "Value : ", value)
	}

}

func SliceOfMap() {

	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Link", "score": "B"},
		{"name": "Mario", "score": "E"},
	}

	for _, student := range students {
		fmt.Println(student["name"], " - ", student["score"])
	}
}
