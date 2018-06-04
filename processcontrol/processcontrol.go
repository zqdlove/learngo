package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(o, p, q float64) float64 {
	if v := math.Pow(o, p); v < q {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, q)
	}
	return q
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 100 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//死循环
	/*
		for {

		}
	*/

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Print("GO run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s\n", os)
	}

	fmt.Println("When is saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Toaday.")
	case today + 1:
		fmt.Println("tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("World.")
	fmt.Println("Hello.") //defer 语句会延迟函数的执行直到上层函数返回

	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //先进后出
	}
	fmt.Println("Done")
}
