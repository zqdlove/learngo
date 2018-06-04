package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needint(x int) int {
	return x*10 + 1
}
func needfloat(f float64) float64 {
	return f * 0.1
}

func main() {
	const f = "%T(%v)\n"

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	var i int
	var b bool
	var fl float64
	var s string

	fmt.Printf("%v %v %v %q\n", i, b, fl, s)

	var x, y int = 3, 4
	var q float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(q)
	fmt.Println(x, y, z)

	v := 5 + 6i
	fmt.Printf("%T\n", v)

	const World = "Rose"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const (
		Truth = true
	)
	fmt.Println("Go rulse?", Truth)

	//fmt.Println(needint(Big))
	fmt.Println(needint(Small))
	fmt.Println(needfloat(Big))
	fmt.Println(needfloat(Small))
}
