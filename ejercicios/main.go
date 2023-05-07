package ejercicios

func main() {

}

func TransformarAPila(q1 *Queue3) Pila4 {
	var p Pila4
	if q1.IsEmpty() {
		return p
	}
	elemento := q1.Dequeue()
	p.Push(elemento)

}

type Pila4 struct {
	Pila []int
}

func (p *Pila4) Push(value int) {
	p.Pila = append(p.Pila, value)

}

type Queue3 struct {
	Cola []int
}

func (q *Queue3) Dequeue() int {
	if len(q.Cola) == 0 {
		return 0
	}
	retorno := q.Cola[0]
	q.Cola = q.Cola[1:len(q.Cola)]
	return retorno
}

func (q *Queue3) IsEmpty() bool {

	return len(q.Cola) == 0
}
