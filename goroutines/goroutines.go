package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {

	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	//	go say("Hello")
	//	say("Rose")
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("%v+%v=%v\n", x, y, x+y)

	c1 := make(chan int, 2)
	c1 <- 1
	c1 <- 2
	fmt.Println(<-c1, <-c1)
	c1 <- 3
	fmt.Println(<-c1)

	c2 := make(chan int, 10)
	go fibonacci1(cap(c2), c2)
	for i := range c2 {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n")

	c3 := make(chan int)
	quit := make(chan int)
	fmt.Println(cap(c3), len(c3))

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%v ", <-c3)
		}
		fmt.Printf("\n")
		quit <- 0
	}()
	fibonacci2(c3, quit)

	tike := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tike:
			fmt.Println("TIKE")
		case <-boom:
			fmt.Println("BOOM")
			return
		default:
			fmt.Println(" .")
			time.Sleep(50 * time.Millisecond)

		}
	}
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Printf("ok ")
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
	close(c)
	close(quit)
}
