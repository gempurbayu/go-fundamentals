package calculation

//nama func kapital artinya public bisa diakses cross package
func Add(number int, numberTwo int) int {
	//return number + numberTwo
	return min(number, numberTwo)
}

//contoh private func yang tidak bisa diakses di package lain
func min(number int, numberTwo int) int {
	return number - numberTwo
}
