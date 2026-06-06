package main
import "os"
import "fmt"

func main() {
	os.Exit(run())
}

func run() int {
	fmt.Println("hello")
	return 0
}