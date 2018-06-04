package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Rose"] = Vertex{
		1.234,
		5.678,
	}
	fmt.Println(m)
	fmt.Println(m["Rose"])

	var n = map[string]Vertex{
		"Chuck": Vertex{
			1.2, 3.4,
		},
		"Zhang": Vertex{
			5.6, 7.8,
		},
	}
	fmt.Println(n)

	var o = map[string]Vertex{
		"Rose":   {9, 8},
		"CHuck:": {7, 6},
	}
	fmt.Println(o)

	p := make(map[string]int)
	p["Chuck"] = 26
	fmt.Println(p, p["Chuck"])
	p["Chuck"] = 27
	fmt.Println(p, p["Chuck"])
	delete(p, "Chuck")
	fmt.Println(p, p["Chuck"])
	v, ok := p["Chuck"]
	fmt.Println("Velue:", v, "Present?", ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)

	}
	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
