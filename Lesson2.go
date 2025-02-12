package main

import "fmt"

// Right
// var someName = "Wassup"

// Wrong
// someName := "Wassup"

func main() {

	//vars
	var nameOne string = "Karim"
	var nameTwo = "NotKarim"
	var nameThree string

	fmt.Println(nameOne, "and", nameTwo, nameThree)

	nameOne = "SomeoneElse"
	nameThree = "NowKarim"

	fmt.Println(nameOne, nameTwo, "and", nameThree)

	nameFour := "WhoAmI"

	fmt.Println(nameFour)

	// int
	var ageOne int = 20
	var ageTwo = 25
	ageThree := 30

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 25 // from -218 to 217
	var numTwo int8 = -128
	var numThree uint = 25  // only >0
	var numFour uint8 = 255 // from 0 to 255

	fmt.Println(numOne, numTwo, numThree, numFour)

	// float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 41239123821749812741294.5
	scoreThree := 1.5 // float64

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
