package ejercicios

// Escriba un método recursivo que tome dos
// números enteros y calcule la multiplicación
// entre ellos, usando sólo sumas
func Multiplicar(a, b int) int {

	// multiplicar un numero por cero es cero
	if b == 0 || a == 0 {
		return 0
	}

	if b < 0 {
		return -(a - Multiplicar(a, b+1))
	}

	if a < 0 {
		return -(b - Multiplicar(b, a+1))
	}

	if b > a {
		return a + Multiplicar(a, b-1)
	}

	return b + Multiplicar(b, a-1)

}
