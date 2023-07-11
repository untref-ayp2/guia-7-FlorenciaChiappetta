package ejercicios

// Escriba un método recursivo que tome un entero n
// y devuelva la suma de los primeros n números naturales.
func Suma(n int) int {
	if n <= 0 {

		return 0
	}
	return n + Suma(n-1)

}
