package main
import "fmt"

func main() {
	age := 25
	name := "John"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("new line! \n")

	// Println
	fmt.Println("Hello, mate!")
	fmt.Println("cheers, mate!")
    fmt.Println("Hello mate, my age is", age, "and my name is", name)

	// Printf (formatted strings) % = format specifier
	fmt.Printf("Hello, my name is %v and my age is %v\n", name, age)
	fmt.Printf("Hello, my name is %q and my age is %q\n", name, age)
	fmt.Printf("Hello, my name is %T and my age is %T\n", name, age)
	fmt.Printf("you scored %f points\n", 225.55)
	fmt.Printf("you scored %0.1f points\n", 225.555555)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("Hello, my name is %v and my age is %v\n", name, age)
	fmt.Println("The saved string is:", str)

}