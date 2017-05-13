package main
import "fmt"
import "time"
func main() {
	// var name = "non"
	// number := 56
	fmt.Println("hello world !!!", validate("hello", 5))
	fmt.Println(time.Now())
}

func validate(input string, number int) int {
	if input == "hello"{
		return 4 * number
	}
	return 0 * number
}