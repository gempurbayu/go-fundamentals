package conditional

import "fmt"

func TestIf() string {
	age := 22
	tingkat := "Balita"

	if age >= 30 {
		tingkat = "Tua"
	} else if age > 25 {
		tingkat = "Dewasa"
	} else if age > 10 {
		tingkat = "Remaja"
	} else {
		tingkat = "Anak Kecil"
	}

	return tingkat
}

func SwitchCase() string {
	number := 2
	var huruf string

	switch number {
	case 1:
		huruf = "satu"
	case 2:
		huruf = "dua"
	case 3:
		huruf = "tiga"
	default:
		huruf = "tidak ada"
	}

	return huruf
}

func ForEx() {

	kata := "golang"
	//example 1
	for i := 1; i <= 3; i++ {
		fmt.Println(kata)
	}

	//example 2
	j := 1
	for j <= 3 {
		fmt.Println(kata)
		j++
	}

	//example 3
	title := "Golang"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index :", index, " Letter : ", string(letter))
		}
	}

	// _ untuk variable penampung, jika terpaksa harus digunakan
	for _, letter := range title {
		fmt.Println(" Letter : ", string(letter))
	}

	for index, letter := range title {
		letterString := string(letter)

		switch letterString {
		case "a", "o":
			fmt.Println("index :", index, " Letter : ", string(letter))
		}
	}

}
