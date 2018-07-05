package main
import(
	"fmt"
//	"os"
	"bufio"
	"strings"
)

func main()  {
	counts := make(map[string]int) 
	s := strings.NewReader("ABC\nDEF\r\nGHI\nJKL\nABC")
	input := bufio.NewScanner(s)
	for input.Scan() {
		counts[input.Text()]++
		
	}

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d %s\n",n, line)
		}		
	}
}