package main

import (
	"fmt"
)

var gv string = "global var"

const gc string = "global constant"

func tryCont() {

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
}
