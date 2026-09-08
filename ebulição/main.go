package main

import "fmt"

// pontoEbulicaoAguaKelvin é a temperatura em que a água ferve, em Kelvin.
const pontoEbulicaoAguaKelvin = 373.0

// kelvinParaCelsius converte uma temperatura de Kelvin para graus Celsius
// usando a fórmula simplificada C = K - 273.
func kelvinParaCelsius(kelvin float64) float64 {
	return kelvin - 273
}

func main() {
	celsius := kelvinParaCelsius(pontoEbulicaoAguaKelvin)

	fmt.Println("Ponto de ebulição da água")
	fmt.Printf("Kelvin:  %.1f K\n", pontoEbulicaoAguaKelvin)
	fmt.Printf("Celsius: %.1f °C\n", celsius)
}