package homework

type ager interface {
	Age() int
}
type Employee struct {
	age int
}
type Customer struct {
	age int
}

func (e *Employee) Age() int {
	return e.age
}

func (e *Customer) Age() int {
	return e.age
}

func older(a ...ager) int {
	res := 0

	for _, e := range a {
		if age := e.Age(); age > res {
			res = age
		}
	}

	return res
}
