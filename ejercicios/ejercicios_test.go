package ejercicios

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEjercicio01(t *testing.T) {
	assert.Equal(t, 1, Suma(1))
	assert.Equal(t, 21, Suma(6))
	assert.Equal(t, 55, Suma(10))
}

func TestEjercicio02(t *testing.T) {
	assert.Equal(t, 2, Factorial(2))
	assert.Equal(t, 24, Factorial(4))
	assert.Equal(t, 120, Factorial(5))
}

func TestEjercicio03(t *testing.T) {
	assert.Equal(t, 0, CantidadDeUnos(0))
	assert.Equal(t, 1, CantidadDeUnos(4))
	assert.Equal(t, 2, CantidadDeUnos(5))
	assert.Equal(t, 3, CantidadDeUnos(42))
}

func TestEjercicio04(t *testing.T) {
	assert.Equal(t, "neuqueN", Invertir("Neuquen"))
	assert.Equal(t, "!odnuM ,aloH", Invertir("Hola, Mundo!"))
}

func TestEjercicio05(t *testing.T) {
	var pila Stack
	pila.Push(1)
	pila.Push(2)
	pila.Push(3)

	assert.Equal(t, 3, CantidadDeElementos(pila))
}

func TestEjercicio06(t *testing.T) {
	assert.Equal(t, 12, MCD(12, 60))
	assert.Equal(t, 12, MCD(60, 12))
	assert.Equal(t, 1, MCD(7, 13))
	assert.Equal(t, 1, MCD(13, 7))
	assert.Equal(t, 2, MCD(2, 4))
	assert.Equal(t, 2, MCD(4, 2))
	assert.Equal(t, 3, MCD(3, 9))
	assert.Equal(t, 3, MCD(9, 3))
}

func TestEjercicio07(t *testing.T) {
	assert.Equal(t, 16, Multiplicar(2, 8))
	assert.Equal(t, 16, Multiplicar(8, 2))
	assert.Equal(t, 0, Multiplicar(0, 8))
	assert.Equal(t, 0, Multiplicar(8, 0))
}

func TestEjercicio08(t *testing.T) {
	cociente, resto := DivisionEntera(8, 2)
	assert.Equal(t, 4, cociente)
	assert.Equal(t, 0, resto)
}

func TestEjercicio09(t *testing.T) {
	assert.Equal(t, 0, SumaArray([]int{}))
	assert.Equal(t, 1, SumaArray([]int{1}))
	assert.Equal(t, 3, SumaArray([]int{1, 2}))
	assert.Equal(t, 6, SumaArray([]int{1, 2, 3}))
}

func TestEjercicio10(t *testing.T) {
	assert.Equal(t, "101010", DecimalABinario(42))
	assert.Equal(t, "10000000000000000000000000000000", DecimalABinario(2147483648))
	assert.Equal(t, "0", DecimalABinario(0))
	assert.Equal(t, "1", DecimalABinario(1))
	assert.Equal(t, "1100", DecimalABinario(12))
}

func TestEjercicio11(t *testing.T) {
	assert.True(t, EsPotencia(8, 2))
	assert.True(t, EsPotencia(1, 2))
	assert.False(t, EsPotencia(7, 2))
}

func TestEjercicio12(t *testing.T) {
	assert.Equal(t, 5, Fibonacci(5))
	assert.Equal(t, 8, Fibonacci(6))
	assert.Equal(t, 13, Fibonacci(7))
	assert.Equal(t, 21, Fibonacci(8))
}

func TestEjercicio13(t *testing.T) {
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 1))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 2))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 3))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 4))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 5))
	assert.False(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 6))
}

func TestEjercicio14(t *testing.T) {
	assert.True(t, SumaElementos([]int{7, 4, 6, 8}, []int{3, 1, 6, 6}, 7))
	assert.False(t, SumaElementos([]int{7, 4, 6, 8}, []int{3, 1, 6, 6}, 2))

	assert.True(t, SumaElementos([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 4))
	assert.False(t, SumaElementos([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 11))
}

func TestEjercicio15(t *testing.T) {
	assert.Equal(t, 4, Pico([]int{1, 2, 3, 4, 5, 4, 3, 2, 1}))
	assert.Equal(t, 8, Pico([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	assert.Equal(t, 0, Pico([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func TestPruebas(t *testing.T) {
	assert.Equal(t, []int{4, 5, 8, 10}, Seleccion([]int{4, 8, 10, 5}))
}
func TestPruebas2(t *testing.T) {
	assert.Equal(t, []int{4, 5, 8, 10}, insercion([]int{4, 8, 10, 5}))
}

func TestPruebas3(t *testing.T) {
	assert.Equal(t, 13, SubsecuenciaSumaMaxima([]int{-2, 11, -4, 2, -3, -10}))
}

func TestPruebas6(t *testing.T) {
	assert.Equal(t, 13, maxSubSum4([]int{-2, 11, -4, 2, -3, -10}))

}

func TestPruebas7(t *testing.T) {
	assert.Equal(t, 13, maxSubSum5([]int{-2, 11, -4, 2, -3, -10}))

}

func TestPruebas8(t *testing.T) {
	assert.Equal(t, 15, SubsecuenciaSumaMaxima([]int{-1, 6, -2, 5, -1, 4, 3, -4, 3}))
}

func TestPruebas9(t *testing.T) {
	assert.Equal(t, 15, maxSubSum4([]int{-1, 6, -2, 5, -1, 4, 3, -4, 3}))

}

func TestPruebas10(t *testing.T) {
	assert.Equal(t, 15, maxSubSum5([]int{-1, 6, -2, 5, -1, 4, 3, -4, 3}))

}

func TestPruebas11(t *testing.T) {
	assert.Equal(t, 7, SubsumaMax([]int{-1, -3, 10, -14}))

}

func TestFunc(t *testing.T) {
	assert.Equal(t, "aloH", InvertirCadena("Hola"))
}

func TestPalindromo(t *testing.T) {
	assert.True(t, Palindromo("NEUQUEN"))
	assert.True(t, Palindromo("BOB"))
	assert.True(t, Palindromo("22522"))
}

// func TestMaximoNumero(t *testing.T) {
// 	assert.Equal(t, 390, MaxNumero(1, 22, 99, 390, -2, 0))
// }

// func TestLinked(t *testing.T) {
// 	list := OrdLinkList.NewLinkedList[int]()
// 	list.Append(1)
// 	list.Append(2)
// 	list.Append(3)
// 	list.Append(5)
// 	fmt.Println(list.String())
// 	list.Insert(4)
// 	fmt.Println(list.String())

// }

func TestCantidadDeElementos(t *testing.T) {
	var pila Stack

	pila.Push(2)
	pila.Push(3)
	pila.Push(4)
	pila.Push(5)
	pila.Push(6)
	assert.Equal(t, 5, CantidadDeElementos(pila))
	fmt.Println(CantidadDeElementos(pila))
}
func TestCantidadDeElementos2(t *testing.T) {
	var pila Stack

	pila.Push(2)
	pila.Push(3)
	pila.Push(4)
	pila.Push(5)
	pila.Push(6)
	pila.Push(5)
	pila.Push(5)
	pila.Push(5)
	pila.Push(5)
	pila.Push(5)
	pila.Push(5)
	pila.Push(5)

	assert.Equal(t, 12, CantidadDeElementos(pila))
	fmt.Println(CantidadDeElementos(pila))

}

func TestMCD(t *testing.T) {
	assert.Equal(t, 10, MCD(50, 120))
	fmt.Println(MCD(50, 120))
	assert.Equal(t, 6, MCD(12, 18))
	fmt.Println(MCD(12, 18))

}

func TestMultiplicar(t *testing.T) {
	assert.Equal(t, 12, Multiplicar(2, 6))
	fmt.Println(Multiplicar(2, 6))

	assert.Equal(t, -12, Multiplicar(-2, 6))
	fmt.Println(Multiplicar(-2, 6))

	assert.Equal(t, -12, Multiplicar(2, -6))
	fmt.Println(Multiplicar(2, -6))

	assert.Equal(t, 12, Multiplicar(-2, -6))
	fmt.Println(Multiplicar(-2, -6))
}

func TestBuscarInfilrado(t *testing.T) {
	prueba := []int{1, 2, 3, 4, 8, 5, 6}
	assert.Equal(t, 8, BuscarInfiltrado(prueba))
	fmt.Println(BuscarInfiltrado(prueba))

	prueba2 := []int{1, 2, 8, 5}
	assert.Equal(t, 8, BuscarInfiltrado(prueba2))
	fmt.Println(BuscarInfiltrado1(prueba2))

	prueba3 := []int{1, 2, 7, 5}
	assert.Equal(t, 7, BuscarInfiltrado(prueba3))
	fmt.Println(BuscarInfiltrado(prueba3))

	prueba4 := []int{1, 2, 3, 4, 5, 6, 9, 7, 8}
	assert.Equal(t, 9, BuscarInfiltrado(prueba4))
	fmt.Println(BuscarInfiltrado(prueba4))

}

func TestBuscarRepetido(t *testing.T) {
	prueba := []int{1, 2, 3, 2, 1, 4, 2}
	assert.True(t, true, BuscarRepetido(prueba))
	fmt.Println(BuscarRepetido(prueba))

	prueba2 := []int{1, 2, 8, 5}
	assert.False(t, false, BuscarRepetido(prueba2))
	fmt.Println(BuscarRepetido(prueba2))

	prueba3 := []int{1, 2}
	assert.False(t, false, BuscarRepetido(prueba3))
	fmt.Println(BuscarRepetido(prueba3))

	prueba4 := []int{3, 2, 3, 4, 3, 3, 9, 7, 8}
	assert.True(t, true, BuscarRepetido(prueba4))
	fmt.Println(BuscarRepetido(prueba4))

	prueba5 := []int{1, 2, 4, 5, 6, 3, 7, 9, 8}
	assert.False(t, false, BuscarRepetido(prueba5))
	fmt.Println(BuscarRepetido(prueba5))

	prueba6 := []int{1, 2, 7, 9, 8}
	assert.False(t, false, BuscarRepetido(prueba6))
	fmt.Println(BuscarRepetido(prueba6))

	prueba7 := []int{1, 1, 1, 8, -3, 5, 7, 7, 9, 1, 7, 6, 2, 2, 4, 7}
	assert.False(t, false, BuscarRepetido(prueba7))
	fmt.Println(BuscarRepetido(prueba7))

	prueba8 := []int{1, 1, 1, 8, -3, 1, 7, 1, 9, 1, 7, 6, 1, 2, 1, 9, 9, 1, 1, 2, 4, 7}
	assert.True(t, true, BuscarRepetido(prueba8))
	fmt.Println(BuscarRepetido(prueba8))

}
func TestHacerA(t *testing.T) {
	fmt.Println(HacerA("ocupacion"))
}
