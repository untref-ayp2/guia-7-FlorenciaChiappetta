package ejercicios

// Escribir un método recursivo que reciba 2 enteros
// n y b y devuelva true si n es potencia de b.
// Por ejemplo: esPotencia(8, 2) devuelve true.
func EsPotencia(n, b int) bool {
   if n == 1 {
        return true
    } else if n < b {
        return false
    } else {
        return esPotencia(n/b, b)
    }
}

func potenciaRecursiva(base float64, exponente int) float64 {
    if exponente == 0 {
        return 1
    } else if exponente < 0 {
        return 1 / potenciaRecursiva(base, -exponente)
    } else {
        return base * potenciaRecursiva(base, exponente-1)
    }
}
