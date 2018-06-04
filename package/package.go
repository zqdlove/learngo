package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}
func del(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func slip(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var m, n int = 2, 3

func main() {
	fmt.Println("My favorite number is ", rand.Intn(1000))
	fmt.Printf("%g\n", math.Nextafter(2, 3))
	fmt.Println(math.Pi)

	fmt.Println(add(7, 9))
	fmt.Println(del(9, 11))
	a, b := swap("Chuck", "Rose")
	fmt.Println(a, b)
	fmt.Println(slip(17))

	var i int
	var c, d, e = true, false, "No"
	f, g, h := "ROSE", true, math.Pi

	fmt.Println(i, c, python, java)
	fmt.Println(m, n, c, d, e)
	fmt.Println(f, g, h)

}
