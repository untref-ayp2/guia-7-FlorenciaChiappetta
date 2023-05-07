package ejercicios

// Escriba un método recursivo que tome una cadena y
// devuelva como resultado la cadena invertida.
func Invertir(cadena string) string {
	if len(cadena) == 0 {
		return cadena
	}
	primeraLetra := string(cadena[0])
	restoDeLaCadena := cadena[1:]

	return Invertir(restoDeLaCadena) + primeraLetra

}
