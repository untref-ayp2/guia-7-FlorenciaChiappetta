package ejercicios

type Conjunto struct {
	mapa map[int]struct{}
}

func (c *Conjunto) Agregar(n int) {
	c.mapa[n] = struct{}{}
}
func (c *Conjunto) Contiene(n int) bool {
	_, ok := c.mapa[n]
	return ok
}

// Sean as y bs dos listas de enteros de tamaño n.
// Escribir una función que reciba como parámetros
// as, bs y un entero x y decida si x se puede
// escribir como suma de un elemento de as más un
// elemento de bs.
func SumaElementos(as, bs []int, x int) bool {
	conjunto := &Conjunto{mapa: make(map[int]struct{})}
	for _, a := range as {
		conjunto.Agregar(a)
	}
	for _, b := range bs {
		if conjunto.Contiene(x - b) {
			return true
		}
	}
	return false
}
