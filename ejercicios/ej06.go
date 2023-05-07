package ejercicios

// Escriba un método recursivo que calcule el máximo
// común divisor entre dos números enteros.
// Nota: Se puede usar el algoritmo de Euclides para
// resolver este problema.
func MCD(a, b int) int {
	if b == 0 {
		return a
	}

	if b > a {
		return MCD(b, a)
	}

	return MCD(b, (a % b))
}
