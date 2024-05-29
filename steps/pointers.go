package steps

import (
	"fmt"
	"runtime"
)

func traceFunction() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Println(frame.Function)
}

// 1. Defining a variable and a pointer to it.
var x = 10
var ptr = &x

func StepOne() {
	traceFunction()
	fmt.Println(x)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}

// 2. We can also manage the memory if we wish.

var ptrTwo = new(int)

func StepTwo() {
	traceFunction()
	fmt.Println(ptrTwo)
	fmt.Println(*ptrTwo)
	*ptrTwo = 10
	ptrTwo = nil
}

// 3. Passing pointers as parameters to functions

func addOne(x *int) {
	*x++
}

func StepThree() {
	traceFunction()
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

func StepFour() {
	traceFunction()
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

func StepFive() {
	traceFunction()
	x := 1
	y := 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
	ptr := double(&x)
	fmt.Println(*ptr)
}
