package try

import (
	"fmt"
)

var GV = "global var"

const GC string = "global constant"

func TryCont() {

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
}
