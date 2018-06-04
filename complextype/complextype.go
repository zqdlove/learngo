package main

import (
	"fmt"
)

type Vertex struct {
	i int
	j int
}

var (
	v1 = Vertex{9, 8}
	v2 = Vertex{i: 7}
	v3 = Vertex{}
	p4 = &Vertex{6, 5}
)

func main() {
	i, j := 7, 9
	p := &i
	fmt.Println(*p)
	*p = 77
	fmt.Println(i)

	p = &j
	*p = *p / 3
	fmt.Println(j)

	fmt.Println(Vertex{5, 6})

	v := Vertex{1, 2}
	v.i = 3
	fmt.Println(v.i, v.j)

	pv := &v
	pv.j = 4
	fmt.Println(pv.i, pv.j)

	fmt.Println(v1, v2, v3, p4)

	var a [2]string
	a[0] = "Hello"
	a[1] = "Rose"
	fmt.Println(a[0], a[1], a)

	p5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(p5)
	for i := 0; i < len(p5); i++ {
		fmt.Printf("p5[%d]=%d\n", i, p5[i])
	}
	fmt.Println(p5[0:9])

	b := make([]int, 5)
	b[4] = 4
	printslice("b", b)
	c := make([]int, 0, 5)
	printslice("c", c)
	d := b[:2]
	printslice("d", d)
	e := d[2:5]
	printslice("e", e)

	var f []int
	fmt.Println(f, len(f), cap(f))
	if f == nil {
		fmt.Println("nil")
	}
	f = append(f, 1)
	printslice("f", f)
	f = append(f, 2)
	printslice("f", f)
	f = append(f, 3, 4, 5, 6, 7)
	printslice("f", f)

	for i, v := range f {
		fmt.Printf("f[%d]=%d\n", i, v)
	}

	g := make([]int, 10)
	for i := range g {
		g[i] = 1 << uint(i)
	}
	for _, value := range g {
		fmt.Printf("%d\n", value)
	}

}
func printslice(s string, b []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(b), cap(b), b)
}
