package main
import "fmt"

func Hello(obj string) string {
	return "Hello " + obj
}

func Goodby(obj string) string {
	return obj
}

func main() {
	fmt.Println(Hello("World"))
}
