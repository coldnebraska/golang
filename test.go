package test

import "fmt"

// both are ints, only type needed on last one
func subtract(x, y int, s1 string) int {
	fmt.Println(s1)
	return x-y
}

func concat(s1 string, s2 string) string {
	return s1+s2
}

func printSubtract(x int, str string) {
	fmt.Println(str, x)
}

func increment(x int) {
	x++
}

func incrementReturn(y int) int{
	y = y+1
	return y
}

func getCoords() (x, y int) { // x and y are initialized with values of zero
	// return -> returns x and y automatically
	return 3, 4
}

func test() {
	fmt.Println("Hello")
	
	height := 75
	if height > 72 {
		fmt.Println("You're over 6ft tall")
	} else if height >60 {
		fmt.Println("You're over 5ft tall")
	} else {
		fmt.Println("You're short")
	}

	// if height := getHeight(75); height > 60 {
	// 	fmt.Println("You're tall")
	// }

	myHeight := 75
	samHeight := 65
	fmt.Println(subtract(myHeight, samHeight, "subtracting..."))

	fmt.Println(concat("My name is ", "David Brush"))

	printSubtract(subtract(10, 7, "subtracting..."), "The value is")

	x := 5
	increment(x)

	fmt.Println("X equals", x)
	// prints 5 since function recieved a copy of x
	// return x in increment function to use new value

	y := 5
	y = incrementReturn(y)
	fmt.Println("Y equals", y)

	// ignore return values
	z, _ := getCoords()
	fmt.Println(z)
}
