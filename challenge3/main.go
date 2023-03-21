package main

import (
	"fmt"
	"strings"
)

func main() {
	var nama string = "Maftuh Ahnan"
	var noReg string = "1955617840-187"
	var title string = "Challenge 3"

	fmt.Println(nama, noReg)
	fmt.Println(title)
	fmt.Println("=====================================")
	greet := "selamat malam"
	split := strings.Split(greet, "")
	splitString(split, 0)
}

func splitString(names []string, loop int) {
	if loop < len(names) {
		fmt.Println(names[loop])
		splitString(names, loop+1)
	} else {
		countAlphabet(names...)
	}
}

func countAlphabet(oke ...string) {

	alph := make(map[string]int)
	for _, v := range oke {
		alph[v] = alph[v] + 1

	}
	fmt.Println(alph)
}
