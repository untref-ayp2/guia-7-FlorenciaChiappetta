package ejercicios

func main() {

}

func EncontrarMaximo(numeros...int)int{

	var pila Pila4
	for _,valores := range numeros{
	pila.Push(valores)
	
	if Pila.Top() <= valores {
		pila.Push(valores)
	}
	
	return pila.Top()
}
}

func TransformarAPila(q1 Queue3) (s1 Pila4) {
	elemento := q1.Dequeue()
	if !q1.IsEmpty() {
	TransformarAPila(q1)
	}
	
	s1.Push(elemento)

	return s1
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

func (q *Queue3) Enqueue(value int) {
	q.Cola = append(q.Cola, value)
}


func (q *Queue3) IsEmpty() bool {

	return len(q.Cola) == 0
}
