package main
import "fmt"

func main() {

	//strings
	var nameOne string = "Manu"
	var nameTwo = "Manu"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "John"
	nameThree = "Jane"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)

	//ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint8 = 25

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 25000000000000000000.5
	scoreThree := 25.98
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}