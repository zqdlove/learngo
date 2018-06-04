package main

import (
	"fmt"
	"math"
	"time"
)

type Myerror struct {
	when time.Time
	what string
}

func (e *Myerror) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &Myerror{
		time.Now(),
		"it did not work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return math.Sqrt(x), nil
	}
	return 0, ErrNegativeSqrt(x)
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
