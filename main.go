package main

import (
	"fmt"

	matchingstring "github.com/sidiqPratomo/golang-hackerrank/5.matching-string"
	arraymanipulation "github.com/sidiqPratomo/golang-hackerrank/6.array-manipulation"
)

func main(){
	stringList := []string{"ab", "ab", "abc"}
    queries := []string{"ab", "abc", "xyz"}

	res := matchingstring.MatchingStrings(stringList,queries)

	fmt.Println("hasil matchingstring = ",res)

	n := int32(7) // Panjang array
    array := [][]int32{
        {1, 2, 200}, // Operasi: Tambah 100 pada elemen-elemen indeks 1 sampai 2 (inklusif)
        {2, 5, 50}, // Operasi: Tambah 100 pada elemen-elemen indeks 2 sampai 5 (inklusif)
        {3, 4, 50}, // Operasi: Tambah 100 pada elemen-elemen indeks 3 sampai 4 (inklusif)
    }

    result := arraymanipulation.ArrayManipulation(n, array)

    fmt.Println("Nilai maksimum setelah operasi:", result)
}