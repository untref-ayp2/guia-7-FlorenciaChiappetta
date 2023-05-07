package ejercicios

// Escriba un método recursivo que calcule Fibonacci de n.
func Fibonacci(numero int) int {
	if numero == 0 {
		return 0
	}
	if numero == 1 {
		return 1
	}
	return Fibonacci(numero-1) + Fibonacci(numero-2)
}
