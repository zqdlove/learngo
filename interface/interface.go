package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}
type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "Hello, Writer\n")

	c := Person{"Chuck Zhang", 26}
	r := Person{"Rose Zhou", 24}
	fmt.Println(c, r)

	addrs := map[string]IPAddr{
		"alibaba": {114, 114, 114, 114},
		"google":  {8, 8, 8, 8},
	}

	for n, a := range addrs {
		fmt.Printf("%v : %v\n", n, a)
	}

}
