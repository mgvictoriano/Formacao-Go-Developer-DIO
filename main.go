package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func main() {
	celsius := 25.0
	fmt.Printf("%.2f°C = %.2f°F\n", celsius, celsiusToFahrenheit(celsius))
	fmt.Printf("%.2f°C = %.2fK\n", celsius, celsiusToKelvin(celsius))
}
