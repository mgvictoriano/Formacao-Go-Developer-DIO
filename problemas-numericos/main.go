package main

import "fmt"

func main() {
	fmt.Println("Numeros entre 1 e 100 divisíveis por 3")

for n := 1; n <= 100; n++ {
	if n%3 == 0 {
		fmt.Println(n)
	}
}
}