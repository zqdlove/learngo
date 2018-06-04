package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Hello struct {
}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello, Chuck")
}

func main() {
	r := strings.NewReader("Hello, ChuckZhang")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v, err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}

}
