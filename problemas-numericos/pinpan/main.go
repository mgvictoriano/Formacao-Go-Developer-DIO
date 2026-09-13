package main

import "fmt"

func main() {

	for n := 1; n <= 100; n++ {

		switch {
		case n%15 == 0:
			fmt.Println("Pin Pan")
		case n%3 == 0:
			fmt.Println("Pin")
		case n%5 == 0:
			fmt.Println("Pan")
		default:
			fmt.Println(n)
		}
	}
}
