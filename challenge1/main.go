package main

import "fmt"

func main() {
	var nama string = "Maftuh Ahnan"
	var noReg string = "1955617840-187"
	var title string = "Challenge 1"
	i := 21
	j := true
	var k float64 = 123.456
	fmt.Println(nama, noReg)
	fmt.Println(title)
	fmt.Println("Nilai variabel i =", i)
	fmt.Printf("Tipe data i = %T \n", i)
	fmt.Print("Menampilkan tanda % \n")
	fmt.Printf("Menampilkan nilai j = %t \n", j)
	fmt.Printf("Menampilkan nilai i = %b \n", i)
	fmt.Printf("Menampilkan unicode rusia = %c \n", '\u042f')
	fmt.Printf("Nilai base 10 = %d \n", i)
	fmt.Printf("Nilai base 8 = %o \n", i)
	fmt.Printf("Nilai base 16 = %x \n", 15)
	fmt.Printf("Nilai base 16 = %X \n", 15)
	fmt.Printf("Menampilkan unicode karakter Я = %U \n", 'Я')
	fmt.Printf("Menampilkan float k = %f \n", k)
	fmt.Printf("Menampilkan float scientific = %E \n", k)
}
