package main
import "fmt"

func Hello(obj string) string {
	return "Hello " + obj
}

func main() {
	fmt.Println(Hello("World"))
}
