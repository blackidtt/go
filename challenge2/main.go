package main

import "fmt"

func main() {
	var nama string = "Maftuh Ahnan"
	var noReg string = "1955617840-187"
	var title string = "Challenge 2"

	fmt.Println(nama, noReg)
	fmt.Println(title)
	fmt.Println("=====================================")
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}
	for j := 0; j <= 10; j++ {
		if j == 5 {
			for x, y := range "САШАРВО" {
				fmt.Printf("character  %#U starts at byte position %d \n", y, x)
			}
		} else {
			fmt.Println("Nilai j = ", j)
		}
	}
}
