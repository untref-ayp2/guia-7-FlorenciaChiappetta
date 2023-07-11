package ejercicios

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// guia1
func SubsumaMax(arreglo []int) int {
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

// guia 2g
// Implementar una cola (QueueS) pero usando internamente dos pilas (stack) para almacenar los datos.
// Analizar el orden de cada uno de los métodos.

type Pila struct {
	Pila []int
}

func (p *Pila) Push(value int) {
	p.Pila = append(p.Pila, value)

}

func (p *Pila) Pop() (int, error) {
	if len(p.Pila) == 0 {
		return 0, errors.New("La pila esta vacia")
	}

	retorno := p.Pila[len(p.Pila)-1]
	p.Pila = p.Pila[:len(p.Pila)-1]

	return retorno, nil
}

func (p *Pila) Top() int {
	if len(p.Pila) == 0 {
		return 0
	}

	return p.Pila[len(p.Pila)-1]

}

func (p *Pila) IsEmpty() bool {
	return len(p.Pila) == 0
}

func (p *Pila) Size() int {
	return len(p.Pila)
}

type Queue struct {
	Cola []any
}

func (q *Queue) Enqueue(value any) {
	q.Cola = append(q.Cola, value)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.Cola) == 0 {
		return nil, errors.New("la cola esta vacia.")
	}
	retorno := q.Cola[0]
	q.Cola = q.Cola[1:len(q.Cola)]
	return retorno, nil
}

func (q *Queue) Front() (any, error) {
	if len(q.Cola) == 0 {
		return nil, errors.New("la cola esta vacia.")
	}
	retorno := q.Cola[0]
	return retorno, nil
}

func (q *Queue) IsEmpty() bool {

	return len(q.Cola) == 0
}

type QueueS struct {
	pila1 *Pila
	pila2 *Pila
}

func (q *QueueS) Enqueue(v int) {

	if q.IsEmpty() {
		q.pila1.Push(v)
	}
	if q.pila2.Size() < q.pila1.Size() {
		q.pila2.Push(v)
	}
	q.pila1.Push(v)
}

func (q *QueueS) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("la cola esta vacia")
	}
	if q.pila1.IsEmpty() {
		q.pila2.Pop()
	}
	result, _ := q.pila1.Pop()

	return result, nil
}

func (q *QueueS) IsEmpty() bool {
	return q.IsEmpty()
}

func (q *QueueS) Front() (any, error) {

	if q.IsEmpty() {
		return nil, errors.New("la cola esta vacia")
	}
	if q.pila1.IsEmpty() {
		retorno := q.pila2.Top()
		return retorno, nil
	}
	retorno := q.pila1.Top()

	return retorno, nil
}

// Escribir una función que que reciba una cadena de caracteres y devuelva la cadena invertida. Analizar el orden.
func InvertirCadena(value string) string {
	var result string
	for _, v := range value {
		result = string(v) + result
	}
	return result
}

// Escribir una función que verifique si una palabra es palíndromo (es decir que una cadena es igual a su inversa.
// Por ejemplo: las cadenas "1456541" y "145541" son palíndromos). Analizar el orden.
func Palindromo(value string) bool {

	if InvertirCadena(value) == value {
		return true
	}

	return false
}

// Escribir una función que evalúe si una cadena de paréntesis, corchetes y llaves está bien balanceada y
// devuelve true si está bien balanceada y false si está mal balanceada.
// Por ejemplo: [()]{}{[()()]()} debe devolver true, mientras que [(]) debe devolver false. Analizar el orden.
// func Balanceada(cadena string) bool {
// 	var pila Pila
// 	simbolosAbiertos := "([{"

// 	for i := 0; i < len(cadena); i++ {
// 		simbolo := string(cadena[i])
// 		if strings.Contains("([{", simbolo) {
// 			pila.Push(simbolo)
// 		} else if !esLaParejitaDe(simbolo, pila.Pop()) {
// 			return false
// 			// {[()]}
// 		}

// 		return pila.IsEmpty()
// 	}
// }

// func esLaParejitaDe(s1, s2 string) {

// 	simbolosAbiertos := "([{"
// 	simbolosCerrados := ")]}"
// 	if strings.Index(s1, simbolosAbiertos) == strings.Index(s2, simbolosCerrados) {
// 		return true
// 	}

// }

// Escribir una función, tal que, dadas dos colas, construya una cola con el resultado de poner una a continuación de la otra.
// Por ejemplo: si q1 = [1,2,3] (con 1 al frente y 3 al final) y q2 = [5,7], el resultado es [1,2,3,5,7] (con 1 al frente y 7 al final). Analizar el orden.
func UnirColas(q1, q2 Queue) (q Queue) {
	if q1.IsEmpty() {
		return q2
	}
	if q2.IsEmpty() {
		return q1
	}
	q1.Enqueue(q2)

	return q1
}

// guia 3

// estructura NODO, cada nodo tiene un valor y un puntero a otro nodo
type Nodo[T any] struct {
	value   T
	puntero *Nodo[T]
}

// Crea un nuevo nodo y me devuelve la dirección del nodo creado.
func newNode[T comparable](value T) *Nodo[T] {

	return &Nodo[T]{value: value, puntero: nil}
}

// estructura linkedlist, cada linked list tiene un nodo cabeza y un nodo cola y una variable tamaño.
type LinkedList[T comparable] struct {
	head   *Nodo[T]
	tail   *Nodo[T]
	tamaño int
}

// Crea una linkedlist y me devuelve la dirección de memoria de la misma.
func NewlinkedList[T comparable]() *LinkedList[T] {

	return &LinkedList[T]{head: nil, tail: nil, tamaño: 0}
}

// Append: Agrega un elemento al final de la lista
func (l1 *LinkedList[T]) Append(value T) {
	newNode := newNode(value)

	if l1.head == nil {
		l1.head = newNode
		l1.tail = newNode
		l1.tail.puntero = nil
		l1.tamaño++
	}

	l1.tail = newNode
	newNode.puntero = nil
	l1.tamaño++
}

// Prepend: Agrega un elemento al principio de la lista
func (l1 *LinkedList[T]) Prepend(value T) {
	newNode := newNode(value)
	if l1.head == nil {
		l1.head = newNode
		l1.tail = newNode
		l1.tail.puntero = nil
		l1.tamaño++
	}
	newNode.puntero = l1.head
	l1.head = newNode
	l1.tamaño++
}

// InsertAt: Agrega un elemento en la posición dada

func (l1 *LinkedList[T]) InsertAt(pos int, value T) {

	if pos < 0 {
		return
	}

	if pos == 0 {
		l1.Prepend(value)
	}
	posActual := l1.head
	for posActual != nil && pos >= 1 {
		posActual = posActual.puntero
		pos--
	}

	nuevoNodo := newNode(value)
	nuevoNodo.puntero = posActual.puntero
	posActual.puntero = nuevoNodo
	l1.tamaño++
}

// Remove: Busca un elemento en la lista y lo elimina

func (l1 *LinkedList[T]) Remove(value T) {

	if l1.head == nil {
		return
	}
	posActual := l1.head
	for posActual != nil {
		if posActual.value == value {
			posActual.puntero = posActual.puntero.puntero
			l1.tamaño--
		}
		posActual = posActual.puntero
	}

}

// Search: Busca un elemento en la lista y me devuelve su posición

func (l1 *LinkedList[T]) Search(value T) int {
	if l1.head == nil {
		return 0
	}
	if l1.head.value == value {
		return 1
	}
	posicion := 2
	current := l1.head.puntero
	for current != nil {
		if current.value == value {
			return posicion
		}
		posicion++
	}
	return posicion
}

// Get: Devuelve el elemento de la lista que se encuentra en una posición dada
func (l1 *LinkedList[T]) Get(pos int) any {
	if l1.head == nil {
		return nil
	}

	posActual := l1.head
	for posActual != nil && pos > 0 {
		posActual = posActual.puntero
		pos--
	}
	return posActual.value
}

// Size: Devuelve la cantidad de elemento que tiene la lista

func (l1 *LinkedList[T]) Size() int {
	return l1.tamaño
}

// String: Devuelve una cadena de caracteres que representa la lista (para usar con Print)
func (l1 *LinkedList[T]) ToString() string {
	var retorno string
	retorno = "["
	posActual := l1.head
	for posActual != nil {
		retorno += fmt.Sprintf("%v", posActual.value)
		if posActual.puntero != nil {
			retorno += " "
		}
		posActual = posActual.puntero
	}
	retorno += "]"
	return retorno
}

// 2- Implementar en la carpeta /slicelist una lista usando slices en lugar de nodos enlazados
// con las mismas operaciones que la lista enlazada. Analizar el orden de cada operación.

// type SliceList[T comparable] struct {
// 	lista []T
// }

// // NewSliceList: Crea una nueva lista vacía
// func NewSliceList[T comparable]() *SliceList[T] {
// 	return &SliceList[T]{}
// }

// // Append: Agrega un elemento al final de la lista

// func (s *SliceList[T]) Append(value T) {
// 	s.lista = append(s.lista, value)
// }

// // Prepend: Agrega un elemento al principio de la lista
// func (s *SliceList[T]) Prepend(value T) {
// 	s.lista = append([]T{value}, s.lista...)
// }

// // InsertAt: Agrega un elemento en la posición dada

// func (s *SliceList[T]) InsertAt(pos int, value T) {
// 	if pos < 0 {
// 		return
// 	}
// 	if pos == 0 {
// 		s.Prepend(value)
// 		return
// 	}
// 	if pos == len(s.lista) {
// 		s.Append(value)
// 		return
// 	}

// 	s.lista = append(s.lista[:pos], s.lista[pos-1:]...)

// 	s.lista[pos-1] = value
// }

// // Remove: Busca un elemento en la lista y lo elimina

// func (s *SliceList[T]) Remove(value T) {

// 	if len(s.lista) == 0 {
// 		return
// 	}
// 	for i := 0; i < len(s.lista); i++ {

// 		if s.lista[i] == value && i == 0 {
// 			s.lista = append(s.lista[1:len(s.lista)])
// 		}

// 		if s.lista[i] == value && i == len(s.lista) {
// 			s.lista = append(s.lista[:len(s.lista)])
// 		}

// 		if s.lista[i] == value {
// 			s.lista = append(s.lista[:len(s.lista)], s.lista[len(s.lista)-1:]...)
// 		}
// 	}

// }

// // String: Devuelve una cadena de caracteres que representa la lista (para usar con Print)
// func (s *SliceList[T]) String() string {
// 	return fmt.Sprintf("%v", s.lista)
// }

// // Get: Devuelve el elemento de la lista que se encuentra en una posición dada

// func (s *SliceList[T]) Get(pos int) T {

// 	return s.lista[pos]
// }

// // Size: Devuelve la cantidad de elemento que tiene la lista

// func (s *SliceList[T]) Size() int {

// 	return len(s.lista)
// }

// // Search: Busca un elemento en la lista y me devuelve su posición sino devuelve -1
// func (s *SliceList[T]) Search(value T) int {

// 	if len(s.lista) == 0 {
// 		return 0
// 	}
// 	for i := 0; i < len(s.lista); i++ {
// 		if s.lista[i] == value {
// 			return i
// 		}

// 		return -1
// 	}

// 	//3- Escribir una función que reciba como parámetros dos listas del mismo tipo y
// 	//devuelva una lista resultante de concatenar la segunda lista al final de la primera

// 	// func (List *LinkedList) UnionDeListas(l1, l2 *LinkedList) *LinkedList[any]{
// 	// 	List = l1
// 	// 	if	len(l2) == 0 {               // REVISAR
// 	// 		return Lista
// 	// 		}
// 	// 	if len(Lista) == 0{
// 	// 			return l2
// 	// 		}

// 	// 	if len(l2) != 0  {
// 	// 	for i := 0; i < len(l2.size); i++{
// 	// 		aux := l2.get(i)
// 	// 		Lista.Append(aux)
// 	// 	}
// 	// 	return Lista
// 	// 	}

// 	// 	return Lista
// 	// }

// 	// // 4- Escribir una función que reciba dos listas del mismo tipo y devuelva la lista que resulta de intercalar uno a uno los elementos de ambas listas.

// 	// func (List *LinkedList) IntercalarListas(l1, l2 *LinkedList) *LinkedList[any]{

// 	// }

// 	// // 5- Implementar una pila y una cola usando la lista enlazada simple

// 	// // func Materias(materias dictionary.Dictionary[string, []string]) dictionary.Dictionary[string, []string] {
// 	// // 	alumnos := dictionary.NewDictionary[string, []string]()

// 	// // 	for _, apellido := range materias.GetKeys() {
// 	// // 		for _, alumno := range materias.Get(apellido) {
// 	// // 			arreglo := []string{}
// 	// // 			if alumnos.Contains(alumno) {
// 	// // 				arreglo = alumnos.Get(alumno)
// 	// // 			}
// 	// // 			arreglo = append(arreglo, apellido)
// 	// // 			alumnos.Put(alumno, arreglo)
// 	// // 		}
// 	// // 	}
// // 	// // 	return alumnos
// // 	// 	// // }
// // }
// 	}

func MaxNumeroOK(elementos ...int) int {
	var pila Pila
	pila.Push(elementos[0])
	for _, numero := range elementos {
		if pila.Top() < numero {
			pila.Push(numero)
		}
	}

	return pila.Top()

}

func MaxNumeroFuncionaOK(elementos []int) int {
	var pila Pila

	if pila.IsEmpty() {
		pila.Push(elementos[0])
	}
	for i := 0; i < len(elementos); i++ {
		if pila.Top() < elementos[i] {
			pila.Push(elementos[i])
		}
	}

	return pila.Top()

}

// //Dada la anterior implementación de una lista enlazada simple, pero que debe mantener el orden de sus elementos.
// //Se pide. Implementar el método Insertar, que reciba como parámetro un elemento y lo inserte en la posición que corresponda. Justificar el orden de Insertar

// func (l *OrdLinkList[int]) Insert(value int) {

// 	NuevoNodo := newNode(value)
// 	if l.head == nil {
// 		l.head = newNode
// 		size++
// 		}
// 	if l.tail.value < value {
// 		 l.tail = NuevoNodo
// 		size++
// 	}

// 	if l.head.value > value {
// 		NuevoNodo.next = l.head
// 		l.head = NuevoNodo
// 		size++
// 		}

// 	posicionActual := l.head
// 	for posicionActual != nil && posicionActual.value < value < posicionActual.next.value{

// 	NuevoNodo.next = posicionActual.next
// 	posicionActual = NuevoNodo
// 	size++

// }

// }

// func (l *OrdLinkList[T]) String() (string){
// 	if l.head == nil {
// 		return "[]"
// 	}
// 	current := l.head
// 	result := "["
// 	for current != nil {
// 		result += fmt.Sprintf("%v", current.value)
// 		if current.next != nil {
// 			result += " "
// 		}
// 		current = current.next
// 	}
// 	result += "]"

// 	return result
// }

// type OrdLinkList[T int] struct {
// 	head *node[int]      // puntero al primer nodo
// 	tail *node[int]      // puntero al último nodo
// 	size int
// }

// func NewOrdLinkList[T int]() *OrdLinkList[T] {
// 	return &OrdLinkList[T]{head: nil, tail: nil, size: 0}
// }

// func (l *OrdLinkList[T]) Append(value int) {
// 	newNode := newNode(value)
// 	if l.head == nil {
// 		l.head = newNode
// 		l.tail = newNode
// 		return
// 	}
// 	l.tail.next = newNode
// 	l.tail = newNode
// 	l.size++
// }

// type node[T int] struct {
// 	value int
// 	next  *node[int]
// }

//	func NewNode(value int) *node[int] {
//		return &node[int]{value: value, next: nil}
//	}
// type Conjunto struct {
// 	mapa map[int]struct{}
// }

// func (c *Conjunto) Agregar(n int) {
// 	c.mapa[n] = struct{}{}
// }

// func (c *Conjunto) Contiene(n int) bool {
// 	_, ok := c.mapa[n]
// 	return ok
// }

//	func SumaElementos2(as, bs []int, x int) bool {
//		conjunto := &Conjunto{mapa: make(map[int]struct{})}
//		for _, a := range as {
//			conjunto.Agregar(a)
//		}
//		for _, b := range bs {
//			if conjunto.Contiene(x - b) {
//				return true
//			}
//		}
//		return false
//	}
//
//	func DIAS() {
//		var diasSemana map[int]string
//		diasSemana = make(map[int]string)
//		diasSemana[1] = "Domingo"
//		diasSemana[2] = "Lunes"
//		diasSemana[3] = "Martes"
//		diasSemana[4] = "Miércoles"
//		diasSemana[5] = "Jueves"
//		diasSemana[6] = "Viernes"
//		diasSemana[7] = "Sábado"
//		//Busca elemento con llave 8
//		string_dia, encontrado := diasSemana[8]
//		if encontrado {
//			fmt.Println("El día 2 es", string_dia)
//		} else {
//			fmt.Println("El día 2 no fue encontrado")
//		}
//		//Eliminar el elemento con llave 2
//		delete(diasSemana, 2)
//		fmt.Println("Tamaño del map:", len(diasSemana))
//	}
func BusquedaBinaria6(datos []int, buscado int) int {
	inicio := 0
	fin := len(datos) - 1
	for inicio <= fin {
		medio := (inicio + fin) / 2
		switch {
		case buscado < datos[medio]:
			fin = medio - 1
		case buscado > datos[medio]:
			inicio = medio + 1
		default:
			return medio
		}
	}
	return -1
}

// Punto 17
// Implementar, por división y conquista, una función que dado un arreglo sin elementos repetidos y casi ordenado
// (todos los elementos se encuentran ordenados, salvo uno), obtenga el elemento fuera de lugar.
//  Indicar y justificar el orden.

func BuscarInfiltrado1(arr []int) int {
	n := len(arr)
	if n <= 2 {
		return arr[0]
	}
	mid := n / 2
	if arr[mid] > arr[mid+1] {
		return arr[mid]
	}
	if arr[mid] < arr[mid-1] {
		return arr[mid+1]
	}
	if arr[mid] > arr[0] {
		return BuscarInfiltrado(arr[mid+1:])
	}

	return BuscarInfiltrado(arr[:mid])
}

func prueba() {
	f, err := os.Create("mi archivo .txt")
	if err != nil {
		panic(err)
	}
	final := 16777215
	for i := 0; i <= final; i++ {
		_, err = f.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}
	}

}

func InsertionSort(arreglo []int) []int {
	for i := 1; i < len(arreglo); i++ {
		actual := arreglo[i]
		j := i - 1
		for j >= 0 && arreglo[j] > actual {
			arreglo[j+1] = arreglo[j]
			i--
		}
		arreglo[j+1] = actual

	}

}

// Set es un conjunto de elementos comparables
type Set[T comparable] struct {
	elementos linkedlist.LinkedList[T]
}

/*********Métodos*********/
// Contains verifica si el elemento pertenece al conjunto
// O(n)
func (s *Set[T]) Contains(element T) bool {
	return s.elementos.Contains(element)
}

// Add agrega un elemento al conjunto
// Si el elemento ya pertenece al conjunto, no hace nada
// O(n) (en el peor de los casos, ya que debe verificar
// si el elemento ya está en el conjunto)
func (s *Set[T]) Add(element T) {
	if !s.Contains(element) { //O(n)
		s.elementos.InsertAt(element, 0) //O(1)
	}
}

// Remove elimina un elemento del conjunto
// Si el elemento no pertenece al conjunto, no hace nada
// O(n)
func (s *Set[T]) Remove(element T) {
	s.elementos.RemoveAt(s.elementos.Search(element))
}

// Size devuelve la cantidad de elementos del conjunto
// O(1)
func (s *Set[T]) Size() int {
	return s.elementos.Size()
}

// String devuelve una representación en cadena del conjunto
// O(n)
func (s *Set[T]) String() string {
	cadena := "Conjunto: {"
	elementos := s.elementos.String()
	inicio := strings.Index(elementos, "[") + 1 // busca el indice donde esta el corchete
	fin := strings.Index(elementos, "]")        // busca el indice donde esta el corchete
	cadena += elementos[inicio:fin]
	cadena += "}"
	return cadena
}

// Values devuelve un arreglo con los elementos del conjunto
func (s *Set[T]) Values() []T {
	var values []T
	for i := 0; i < s.Size(); i++ {
		values = append(values, s.elementos.Get(i))
	}
	return values
}

/*********Operaciones sobre conjuntos *********/

// NewSet crea un nuevo conjunto con los elementos que recibe como parámetros
// Si no recibe ningún elemento, crea un conjunto vacío
// O(n) donde n es la cantidad de elementos

// Dada la implementación de Set vista en clase (implementado sobre una lista enlazada simple) que tiene los métodos:
// Contains, Add, Remove, Size, String y Values y soporta las siguientes operaciones que en todos los casos devuelven un conjunto nuevo:
// NewSet, Union, Intersection, Difference, Subset y Equal.
//
//	Se pide reescribir la operación Union para que reciba como parámetro un número indefinido de conjuntos y devuelva la Union de todos ellos.
func NewSet[T comparable](elementos ...T) *Set[T] {
	conjunto := &Set[T]{*linkedlist.NewLinkedList[T]()}
	for _, elemento := range elementos {
		conjunto.Add(elemento)
	}
	return conjunto
}

// Union devuelve un nuevo conjunto con la unión de los conjuntos recibidos
// O(n*m), donde n y m son los tamaños de los conjuntos s1 y s2 respectivamente
func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for i := 0; i < s1.Size(); i++ {
		result.Add(s1.elementos.Get(i))
	}
	for i := 0; i < s2.Size(); i++ {
		result.Add(s2.elementos.Get(i))
	}
	return result
}

// Intersection devuelve un nuevo conjunto con la intersección de los conjuntos
// recibidos
// O(n*m) donde n y m son los tamaños de los conjuntos s1 y s2 respectivamente
func Intersection[T comparable](s1, s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for i := 0; i < s1.Size(); i++ {
		elemento := s1.elementos.Get(i)
		if s2.Contains(elemento) {
			result.Add(elemento)
		}
	}
	return result
}

// Difference devuelve un nuevo conjunto con la diferencia de los conjuntos
// recibidos
// La diferencia entre s1 y s2, son todos los elementos de s1 que no están en s2
// O(n*m) donde n y m son los tamaños de los conjuntos s1 y s2 respectivamente
func Difference[T comparable](s1, s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for i := 0; i < s1.Size(); i++ {
		elemento := s1.elementos.Get(i)
		if !s2.Contains(elemento) {
			result.Add(elemento)
		}
	}
	return result
}

// Subset verifica si el conjunto s2 es subconjunto del conjunto s1
// O(n*m) donde n y m son los tamaños de los conjuntos s1 y s2 respectivamente
func Subset[T comparable](s1, s2 *Set[T]) bool {
	for i := 0; i < s2.Size(); i++ {
		if !s1.Contains(s2.elementos.Get(i)) { //valida elemento por elemento si existe en el conjunto1 sino, retorna falso
			return false
		}
	}
	return true
}

// Equal verifica si los conjuntos recibidos son iguales
// O(n^2) donde n es el tamaño de los conjuntos s1 y s2
func Equal[T comparable](s1, s2 *Set[T]) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

// PARCIAL PUNTO 1

func UnionParcial[T comparable](Conjuntos ...*Set[T]) *Set[T] {
	result := NewSet[T]()
	for _, conjunto := range Conjuntos {
		valores := conjunto.Values()
		for _, valor := range valores {
			result.Add(valor)
		}
	}

	return result
}

// Diccionario es un conjunto de entradas formadas por pares únicos (clave: valor)
type Dictionary[K comparable, V any] struct {
	mapa map[K]V
}

/*********Métodos*********/
// Crea un diccionario
// O(1)
func NewDictionary[K comparable, V any]() Dictionary[K, V] {
	dict := Dictionary[K, V]{mapa: make(map[K]V)}
	return dict
}

// Agrega una nueva entrada al diccionario, si existe pisa el valor para esa entrada
// O(1)
func (dict *Dictionary[K, V]) Put(key K, val V) {
	dict.mapa[key] = val
}

// Remueve una entrada del diccionario
// O(1)
func (dict *Dictionary[K, V]) Remove(key K) bool {
	var exists bool
	_, exists = dict.mapa[key]
	if exists {
		delete(dict.mapa, key)
	}
	return exists
}

// Verifica si existe una entrada para ese valor de clave
// O(1)
func (dict *Dictionary[K, V]) Contains(key K) bool {
	var exists bool
	_, exists = dict.mapa[key]
	return exists
}

// Devuelve el valor para esa clave
// O(1)
func (dict *Dictionary[K, V]) Get(key K) V {
	return dict.mapa[key]
}

// Devuelve la cantidad de entradas del diccionario
// O(1)
func (dict *Dictionary[K, V]) Size() int {
	return len(dict.mapa)
}

// Devuelve un arreglo con las claves del diccionario
// O(n)
func (dict *Dictionary[K, V]) GetKeys() []K {
	var dictKeys []K
	dictKeys = []K{}
	var key K
	for key = range dict.mapa {
		dictKeys = append(dictKeys, key)
	}
	return dictKeys
}

// Devuelve un arreglo con las valores del diccionario
// O(n)
func (dict *Dictionary[K, V]) GetValues() []V {
	var dictValues []V
	dictValues = []V{}
	var key K
	for key = range dict.mapa {
		dictValues = append(dictValues, dict.mapa[key])
	}
	return dictValues
}

// PARCIAL PUNTO 2

// Escribir una función que recibe un diccionario cuyas claves son los nombres de los alumnos y
// como valor una lista con los días que asistieron a clase. Debe devolver un diccionario con clave fecha
// y valor la lista de alumnos que asistieron en dicha fecha.  Por ejemplo, si la entrada es

// {"Ana": [ "Mie 10", "Vie 12"], "Luz": [ "Vie 12",  "Mie 17"],  "Pedro": ["Mie  10", "Mie 17"]},
// debe devolver  {“Mie 10”: [“Ana", "Pedro"]), “Vie 12”: [“Ana", "Luz”], “Mie 17”: [“Luz”, "Pedro"]}

func EjercicioDiccionario(entrada Dictionary[string, []string]) Dictionary[string, []string] {
	alumnos := entrada.GetKeys()
	aux := NewDictionary[string, []string]()
	for _, alumno := range alumnos {
		dias := entrada[alumno]
		for _, dia := range dias {
			if !aux.Contains(dia) {
				aux.Put(dia, []string{})
			}
			aux[dia] = append(aux[dia], alumno)
		}
	}
	return aux

}
