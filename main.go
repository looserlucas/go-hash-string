package main

import "fmt"

var hash, pow []int64
var p int64 = 31
var m int64 = 1e9 + 9

func getHash(i, j int) int64 {
	return (hash[j] - hash[i-1]*pow[j-i+1] + m*m) % m
}

func main() {
	s := "eqwjklavnklas"
	pt := "a"
	achar := "a"

	var hashP int64
	lenT := len(s)
	lenP := len(pt)
	s = " " + s
	pt = " " + pt

	pow = make([]int64, lenT+1)
	hash = make([]int64, lenT+1)
	pow[0] = 1
	for i := 1; i <= lenT; i++ {
		pow[i] = (pow[i-1] * p) % m
		hash[i] = (hash[i-1]*p + int64(s[i]) - int64(achar[0]) + int64(1)) % m
	}

	hashP = 0
	for i := 1; i <= lenP; i++ {
		hashP = (hashP*p + int64(pt[i]) - int64(achar[0]) + int64(1)) % m
	}
	fmt.Println(hashP)

	for i := 1; i <= lenT-lenP+1; i++ {
		if hashP == getHash(i, i+lenP-1) {
			fmt.Println(i)
		}
	}
}
