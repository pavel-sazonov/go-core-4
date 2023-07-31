package main

import (
	"io"
	"os"
)

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

func main() {
	var i any
	a := []any{"задание3", 8, false, 0.4, Employee{19}}
	a1 := []any{8, false, 0.4, Employee{19}, i}
	a2 := make([]any, 0, 1)
	writeStr(os.Stdout, a...)
	writeStr(os.Stdout, a1...)
	writeStr(os.Stdout, a2...)
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

// задание 3: вывод в writer только строк
func writeStr(w io.Writer, t ...any) {
	for _, obj := range t {
		if s, ok := obj.(string); ok {
			w.Write([]byte(s))
		}
	}
}
