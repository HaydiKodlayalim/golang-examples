package main

import "fmt"

func main() {
	ilkFonsiyonOrnegim()
	fmt.Println("Merhaba", 1234, "Metin", 34)

	adimAt(30, 45, "sag")

	adimAt(30, 45, "sol")

	adimAt(45, 45, "sag")

	adimAt(30, 45, "sol")

	metodSonucu := kemikOradaMi()

	fmt.Println(metodSonucu)

	fmt.Println(adimAtYeniFonk(3, 5, "sol"))
}

func ilkFonsiyonOrnegim() {
	fmt.Println("ilkFonsiyonOrnegim metodu calisti")
}

func adimAt(yukari int, ileri int, ayak string) {
	fmt.Println("ayak:", ayak, "yukari:", yukari, "ileri:", ileri)
}

func kemikOradaMi() string {
	return "Kemik burada"
}

func adimAtYeniFonk(yukari int, ileri int, ayak string) string {
	fmt.Println("Yeni Fonk: ", "ayak:", ayak, "yukari:", yukari, "ileri:", ileri)
	return "Deger"
}

// string
// int -500 , 500
// float 34.56

/*
1 - basla
2 - sag 30Dere yukari 45 D ileri
3 - ayagi indir
4 - sol 30Dere yukari 45 D ileri
5 - ayagi indir
6 - sag 30Dere yukari 45 D ileri
7 - ayagi indir
8 - kemik orada mi
9 -
*/
