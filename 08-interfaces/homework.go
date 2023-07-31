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

// задание 1: возвращает старший возраст
func older(a ...ager) int {
	res := 0

	for _, e := range a {
		if age := e.Age(); age > res {
			res = age
		}
	}

	return res
}

// задание 2: возвращает объект с самым старшим возрастом
func olderObj(a ...ager) ager {
	age := 0
	var person ager

	for _, p := range a {
		switch t := p.(type) {
		case *Employee:
			if t.age > age {
				age = t.age
				person = p
			}
		case *Customer:
			if t.age > age {
				age = t.age
				person = p
			}
		}
	}
	return person
}
