package main

import (
	"fmt"

	dsaCore "github.com/kroncatti/go-playground/dsa/core"
)

func main() {
	maxVal := dsaCore.MaxValue([]int{1, 2, 10, 3, 42})
	fmt.Println(maxVal)
}
