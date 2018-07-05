package main
import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	counts:=make(map[string]int)
	files:=os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}else{
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n",err)
				continue
			}
			countLines(f,counts)
			for line, n := range counts{
				if n > 1 {
					fmt.Printf("%s\t%d\t%s\n",f.Name(), n, line)
					
				}
			}
		
			f.Close()
		}
	}
	

}

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
	}
}