package ejercicios

// Escribir un método que calcule recursivamente cuántos
// elementos hay en una pila y no altere el contenido de
// la misma. La pila sólo tiene los métodos Push, Pop y isEmpty.
func CantidadDeElementos(pila Stack) int {
	var n int
	if pila.IsEmpty() {
		return 0
	}

	pila.Pop()
	n++

	return n + CantidadDeElementos(pila)
}
