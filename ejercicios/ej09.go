package ejercicios

// Escribir un método recursivo que tome un
// array de números enteros y devuelva la suma
// de todos sus elementos
func SumaArray(v []int) int {
	if len(v) == 0 {
		return 0
	}

	return v[0] + SumaArray(v[1:])
}
