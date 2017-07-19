// https://gobyexample.com/channel-directions

package try

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg + " and I got it now"
}

func TryChannels() {
	ch := make(chan string, 5)
	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"
	ch <- "5"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message by sender")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
