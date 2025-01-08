package main

import (
	"calculate-api/pkg/calc"
	"fmt"
)

func main() {
	var expr string
	_, _ = fmt.Scanln(&expr)
	res, err := calc.Calc(expr)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
