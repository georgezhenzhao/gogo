// originally from https://gobyexample.com/pointers

package try

import "fmt"

func zeroval(p int) {
	p = 0
}

func zeroptr(p *int) {
	*p = 0
}

// test pointers
func TryPointer() {

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
