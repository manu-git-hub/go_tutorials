package main
import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	//classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "luigi", "peach", "bowser"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, value := range names {
		fmt.Printf("the value is %v\n", value)
		value = "new value"
	}

	fmt.Println("loop is over")
	fmt.Println(names)
}