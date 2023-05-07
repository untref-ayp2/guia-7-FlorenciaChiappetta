package ejercicios

func Seleccion(items []int) []int {

	for i := 0; i < len(items)-1; i++ {

		valorMax := items[i]
		pos := i

		for j := i + 1; j < len(items); j++ {

			if valorMax > items[j] {
				valorMax = items[j]
				pos = j
			}

		}
		items[i], items[pos] = valorMax, items[i]

	}
	return items
}

func insercion(items []int) []int {

	for i := 0; i < len(items); i++ {
		value := items[i]
		j := i - 1
		for j >= 0 && items[j] > value {
			items[j+1] = items[j]
			j = j - 1
		}
		items[j+1] = value
	}
	return items
}

// Dado un arreglo devuelve la posicion inicial, la posición final y el valor
// de la subsecuencia cuya suma es máxima dentro del arreglo original
func SubsecuenciaSumaMaxima(arreglo []int) int {
	sumaMaxima := 0
	for i := 0; i < len(arreglo); i++ {
		sumaLocal := 0
		for j := i; j < len(arreglo); j++ {
			sumaLocal += arreglo[j]
			if sumaLocal > sumaMaxima {
				sumaMaxima = sumaLocal
			}
		}
	}
	return sumaMaxima
}

func maxSubSum4(arreglo []int) int {
	sumaMax, suma := 0, 0
	for i := 0; i < len(arreglo); i++ {
		suma += arreglo[i]
		if suma > sumaMax {
			sumaMax = suma
		}
		if suma < 0 {
			suma = 0
		}
	}
	return sumaMax
}

func maxSubSum5(arreglo []int) int {
	maxSum := 0
	for i := 0; i < len(arreglo); i++ {
		thisSum := 0
		for j := i; j < len(arreglo); j++ {
			thisSum += arreglo[j]
			if thisSum > maxSum {
				maxSum = thisSum
			}
		}

	}
	return maxSum
}

// func CreateHash(byteStr []byte) []byte {
// 	var hashVal hash.Hash
// 	hashVal = sha1.New()
// 		hashVal.Write(byteStr)

// 	var bytes []byte

// 	bytes = hashVal.Sum(nil)
// 	return bytes
// }
