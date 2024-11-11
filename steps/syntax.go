package steps

import (
	"fmt"

	"github.com/kroncatti/go-playground/utils"
)

// 1. Variables

func syntaxOne() {
	utils.TraceFunction()
	mileage, company := 80276, "Tesla"

	fmt.Println(mileage, company)
}

func Syntax() {
	syntaxOne()
}
