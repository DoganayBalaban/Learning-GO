package main

import "fmt"

type shape interface {
	area() float64
}

type triangle struct {
	a float64
	h float64
}

func (t triangle) area() float64 {
	return (t.a * t.h) / 2
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return s.a * s.a
}

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func areaFind(shapes ...shape) {
	for _, shape := range shapes {
		fmt.Println("Area : ", shape.area())
	}
}
func main() {
	t := triangle{5, 10}
	s := square{4}
	r := rectangle{5, 7}

	areaFind(t, s, r)

}
