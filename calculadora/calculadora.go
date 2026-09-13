package calculadora

import "errors"

var ErrDivisaoPorZero = errors.New("não é possível dividir por zero")

func Somar(a, b float64) float64 {
	return a + b
}

func Subtrair(a, b float64) float64 {
	return a - b
}

func Multiplicar(a, b float64) float64 {
	return a * b
}

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisaoPorZero
	}
	return a / b, nil
}

