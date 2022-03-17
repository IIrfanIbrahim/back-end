package main

import (
	"fmt"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {
	// TODO: answer here
	var j int
	var t int
	fmt.Print("Masukkan jari-jari alas : ")
	fmt.Scan(&j)
	fmt.Print("\nMasukkan tinggi alas : ")
	fmt.Scan(&t)

	result := j + t
	fmt.Println(result)
}
