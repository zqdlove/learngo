package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
type Abser interface {
	Abs() float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) Scala(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Myfloat float64

func (m Myfloat) Abs() float64 {
	if m > 0 {
		return float64(m)
	}
	return float64(-m)
}

func main() {

	v := &Vertex{3, 4}

	fmt.Println(v.Abs())
	v.Scala(5)
	fmt.Println(v.Abs())

	m := Myfloat(-math.Sqrt2)
	fmt.Println(m.Abs())

	var a Abser
	a = &Vertex{6, 8}
	fmt.Println(a.Abs())
	a = m
	fmt.Println(a.Abs())

}
