package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
Go's basic types are

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128
The example shows variables of several types, and also that variable declarations may be
"factored" into blocks, as with import statements.
The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems
and 64 bits wide on 64-bit systems. When you need an integer value you should use
int unless you have a specific reason to use a sized or unsigned integer typ
*/
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables declared with initial value  or default value
	var idNumber int
	var floatNumber float64
	var isBoolean bool
	var emptyString string
	fmt.Printf("idNumber ={%v} floatNumber= {%v} isBoolean= {%v} emptyString={%q}\n",
		idNumber, floatNumber, isBoolean, emptyString)

	// Type conversions
	intValue := 47
	floatValue := float64(intValue)
	uintValue := uint(floatValue)
	fmt.Printf("intValue ={%v} floatValue= {%v} uintValue= {%v} \n",
		intValue, floatValue, uintValue)
	var a, b = 2, 4
	var float64Value float64 = math.Sqrt(float64(a + b)) // here also explicit type cast required
	fmt.Println("float64Value", float64Value)

	var float64Value2 float64 = math.Sqrt((2 + 4))
	fmt.Println("float64Value2", float64Value2)
}
