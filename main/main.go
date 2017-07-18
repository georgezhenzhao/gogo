package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("\n********************")
	if len(os.Args) > 1 {

		args := os.Args[1:]
		fmt.Println("gogo is starting with params given:")
		fmt.Println(args)

	} else {

		fmt.Println("gogo is starting without params given.")
		fmt.Println("Nothing we can do now.")
		os.Exit(1)
	}
	fmt.Println("********************\n")

	if os.Args[1] == "try" {
		tryFunc()
		tryVar()
		tryCont()
		os.Exit(1)
	}

	if os.Args[1] == "g" {

		fmt.Println(gv)
		fmt.Println(gc)
		os.Exit(1)

	}

	fmt.Println("Please enter 'q' to exit.")

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "q" {
		fmt.Print("Enter your text: ")
		scanner.Scan()
		text = scanner.Text()
		if text != "q" {
			fmt.Println("Your text was: ", text)
		}
	}

}
