package main

import (
	"fmt" // package names are in double quote separated by space not with comma
	"math"
	"math/rand"
	"reflect"
	"time"
)

var goBasic string

func main() {
	goBasic = "This program is about basics of Go programming language"
	fmt.Println("Hello from Go language  ->", goBasic)
	// use of rand package
	fmt.Println("Your luck number is  ", rand.Int())

	// use of math package
	fmt.Println("Square root of 3 is ", math.Sqrt(3))

	// A name is exported in Go if it begins with Capital letter

	fmt.Println("Exported in go for different package  ", math.Ln2)

	// simple add function
	fmt.Println("Multiply  of  two number is ", multiply(4, 5))

	// Function to return two argument
	x, y := swapNames("Java", "Go")             // first argument value is Java and second argument value is Go
	fmt.Println("After swap x and y is ", x, y) // should be X=Go and Y=Java

	// "Naked" return, A return statement without arguments returns the named return values
	reminder, multiplied := divideAndMultiplyByTwo(15)
	fmt.Println("Divided by 2 ", reminder, "  Multiplied by 2", multiplied)

	// variable declaration here multiple/list of variables are declared at once of type sting
	// var can be used at package level also
	var itemName, itemCode string
	itemName = "Test item"
	itemCode = "Test code"
	fmt.Println("itemName ", itemName, " itemCode ", itemCode)

	// A var declaration can include initializers, one per variable.
	var i, j int = 1, 2 // here type declaration can be removed
	var c, python, java = true, false, "no!"
	var num = 200 // data type is decided by the compiler based on the value assigned
	fmt.Println(i, j, c, python, java, num)

	// Short variable declarations  using type inference

	today := time.Now() // here type declaration is done base on the initialized value
	fmt.Println("Today's date is ", today, "today is of type ", reflect.TypeOf(today))

	amount := 107.99
	fmt.Println("amount is of type ", reflect.TypeOf(amount))
	// A function can return n number of values
	//var p, q, r, s, t int = return5Value()
	p, q, r, s, t := return5Value()
	//or var p, q, r, s, t int = return5Value()
	fmt.Printf("p =%v q =%v r =%v s =%v t =%v ", p, q, r, s, t)

}

func swapNames(a string, b string) (string, string) {
	return b, a
}

// func multiply(x int, y int) int
func multiply(x, y int) int {
	return x * y
}

// In naked return the argument name must be declared like here (reminder int, multiplied int)
func divideAndMultiplyByTwo(x int) (reminder int, multiplied int) {
	reminder = x / 2
	multiplied = x * 2
	return
}

func return5Value() (a int, b int, c int, d int, e int) {
	return 1, 3, 4, 5, 6
}
