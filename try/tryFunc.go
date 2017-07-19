package try

import (
	"fmt"
)

func TryFunc() {

	fmt.Println("tryFunc is starting.")

	resDisplay := display("go", "lang")
	fmt.Println(resDisplay)

	fmt.Println(plus(1, 1))

	if num := 1; num < 0 {
		fmt.Println(num, "is less than 0")
	} else if num < 2 {
		fmt.Println(num, "is a digit")
	} else {
		fmt.Println(num, "others")
	}

	fmt.Println("tryFunc is completed.")

}

func plus(a int, b int) (int, int) {

	return a + b, a - b
}

func display(a string, b string) string {

	return a + b
}
