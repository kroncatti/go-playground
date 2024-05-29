package steps

import (
	"fmt"

	"github.com/kroncatti/go-playground/utils"
)

// 1. Defining a variable and a pointer to it.
var x = 10
var ptr = &x

func stepOne() {
	utils.TraceFunction()
	fmt.Println(x)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}

// 2. We can also manage the memory if we wish.

var ptrTwo = new(int)

func stepTwo() {
	utils.TraceFunction()
	fmt.Println(ptrTwo)
	fmt.Println(*ptrTwo)
	*ptrTwo = 10
	ptrTwo = nil
}

// 3. Passing pointers as parameters to functions

func addOne(x *int) {
	*x++
}

func stepThree() {
	utils.TraceFunction()
	x := 10
	fmt.Println(x)
	addOne(&x)
	// This changes the global variable
	fmt.Println(x)
}

// 4. Working with arrays and pointers

func printArray(arr *[4]int) {
	fmt.Println(*arr)
}

func stepFour() {
	utils.TraceFunction()
	arr := [4]int{1, 2, 3, 4}
	ptr := &arr
	fmt.Println(*ptr)
	printArray(ptr)
}

// 5. Working with functions and pointers

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func double(x *int) *int {
	result := *x * 2
	return &result
}

func stepFive() {
	utils.TraceFunction()
	x := 1
	y := 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
	ptr := double(&x)
	fmt.Println(*ptr)
}

func Pointers() {
	stepOne()
	stepTwo()
	stepThree()
	stepFour()
	stepFive()
}
