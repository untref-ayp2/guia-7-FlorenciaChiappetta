package ejercicios

import "strings"

// Punto 17
// Implementar, por división y conquista, una función que dado un arreglo sin elementos repetidos y casi ordenado
// (todos los elementos se encuentran ordenados, salvo uno), obtenga el elemento fuera de lugar.
//  Indicar y justificar el orden.

func BuscarInfiltrado(arr []int) int {
	n := len(arr)
	if n <= 2 {
		return arr[0]
	}
	mitad := n / 2
	if arr[mitad] > arr[mitad+1] {
		return arr[mitad]
	}
	if arr[mitad] < arr[mitad-1] {
		return arr[mitad-1]
	}
	if arr[mitad] > arr[0] {
		return BuscarInfiltrado(arr[mitad+1:])
	}

	return BuscarInfiltrado(arr[:mitad])
}

// Punto 16
// Implementar una función (que utilice división y conquista) que dado un arreglo de n números enteros devuelva true o false
// según si existe algún elemento que aparezca más de la mitad de las veces. Justificar el orden de la solución.

func BuscarRepetido(arreglo []int) bool {
	contador := 1

	if len(arreglo) == 0 {
		return false
	}

	if len(arreglo) == 1 {
		return true
	}
	mitad := len(arreglo) / 2
	if arreglo[mitad] == arreglo[0] {
		contador++
		return BuscarRepetido(arreglo[1:mitad])
	}
	BuscarRepetido(arreglo[:mitad])

	if arreglo[mitad] == arreglo[len(arreglo)-1] {
		contador++
		return BuscarRepetido(arreglo[mitad+1:])
	}
	BuscarRepetido(arreglo[mitad+1:])

	if contador >= mitad {
		return true
	}
	return false
}

// Punto 15
// Escribir una función que reciba un arreglo de enteros de tamaño N, ordenado ascendente y sin repetidos y determine si el arreglo es mágico.
//  El orden de la función debe ser O(log(N)). Un arreglo es mágico si existe alguna posición i dentro del mismo, tal que a[i] = i.
//Por ejemplo [-7, 1, 0, 3, 7, 10] es mágico porque a[3] = 3. Mientras que [1, 2, 3, 4, 5, 6, 7] no lo es. Justificar el orden.

func ArregloMagico(arreglo []int) bool {
	if len(arreglo) == 0 {
		return false
	}

	if arreglo[0] == 0 || arreglo[len(arreglo)-1] == len(arreglo)-1 {
		return true
	}
	medio := len(arreglo) / 2

	if arreglo[medio] == medio {
		return true
	}
	ArregloMagico(arreglo[:medio])
	ArregloMagico(arreglo[medio+1:])

	return false
}

// Punto 14
// Escribir una función que devuelva el elemento máximo y mínimo de un arreglo de enteros utilizando la estrategia de dividir y conquistar.
// Calcular el orden de la solución propuesta. Comparar con una solución que no use DyC

func MaxyMin(arreglo []int) (int, int) {
	if len(arreglo) == 0 {
		return 0, 0
	}
	if len(arreglo) == 1 {
		return arreglo[0], arreglo[0]
	}

	var min int
	var max int
	medio := len(arreglo) / 2

	if arreglo[0] < min {
		min = arreglo[0]
	}
	if arreglo[0] > max {
		max = arreglo[0]
	}
	if arreglo[len(arreglo)-1] < min {
		min = arreglo[len(arreglo)-1]
	}
	if arreglo[len(arreglo)-1] > max {
		max = arreglo[len(arreglo)-1]
	}

	if arreglo[medio] < min {
		min = arreglo[medio]
	}
	if arreglo[medio] > max {
		max = arreglo[medio]
	}

	MaxyMin(arreglo[:medio])
	MaxyMin(arreglo[medio+1:])

	return max, min
}

//Punto 13
func HacerA(palabra string) string {
	var i int
	if i == len(palabra) {
		return palabra
	}
	aux := string(palabra[i])
	vocalesMinuscula := "eiou"
	vocalesMayuscula := "EIOU"

	if strings.Contains(aux, vocalesMinuscula) {
		aux = "a"
	}
	if strings.Contains(aux, vocalesMayuscula) {
		aux = "A"
	}
	i++
	return HacerA(palabra[:len(palabra)/2]) + HacerA(palabra[len(palabra)/2:])

}

//Punto 9

// func main() {
// 	m := newMapaDeBits()
// 	dias := []int{11, 14, 19, 21, 30}
// 	m.RegistroDeLluvias("enero", dias)
// 	fmt.Println(m.BuscarLluviasEn("enero"))

// }

type Dias struct {
	Dias []int
}

type MapaDeBits[s string, d Dias] struct {
	mapa map[string]Dias
}

func newMapaDeBits[s string, d Dias]() MapaDeBits[s, d] {
	m := MapaDeBits[s, d]{make(map[string]Dias)}
	return m
}

func (m *MapaDeBits[k, v]) RegistroDeLluvias(mes string, dias []int) {
	m.mapa[mes] = Dias{Dias: dias}
}

func (m *MapaDeBits[k, v]) BuscarLluviasEn(mes string) Dias {

	return m.mapa[mes]
}

func (m *MapaDeBits[s, d]) LluviasDe(mes1, mes2 string) []int {
	dias := []int{}

	dias = append(dias, m.mapa[mes1].Dias...)
	dias = append(dias, m.mapa[mes2].Dias...)

	return dias
}

//Punto 2

func EncontrarMaximoConPila(numeros []int) int {

	var pila Pila6
	for _, valores := range numeros {
		if pila.Top() <= valores {
			pila.Push(valores)
		}
	}

	return pila.Top()
}

type Pila6 struct {
	Pila []int
}

func (p *Pila6) Push(value int) {
	p.Pila = append(p.Pila, value)

}

func (p *Pila6) Top() int {
	if len(p.Pila) == 0 {
		return 0
	}

	return p.Pila[len(p.Pila)-1]

}
