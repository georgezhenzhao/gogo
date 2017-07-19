package main

import (
	"bufio"
	"fmt"
	"os"

	"../utils"
)

func main() {

	utils.Initlog()

	//utils.Trace.Println("Trace log")
	utils.Info.Println("-- gogo is a GoLang project --")
	//utils.Warning.Println("Warning log")
	//utils.Error.Println("Error log")

	fmt.Println("\n********************")
	if len(os.Args) > 1 {

		args := os.Args[1:]
		fmt.Println("gogo is starting with params given:")
		fmt.Println(args)
		utils.Info.Println("gogo is starting with params given")

	} else {

		fmt.Println("gogo is starting without params given.")
		fmt.Println("Nothing we can do now.")
		utils.Info.Println("Nothing we can do now.")
		os.Exit(1)
	}
	fmt.Println("********************\n")

	if os.Args[1] == "test" {
		tryFunc()
		tryVar()
		tryCont()

		fmt.Println(gv)
		fmt.Println(gc)
		os.Exit(1)

	} else if os.Args[1] == "run" {
		fmt.Println("Please enter 'q' to exit.")
		fmt.Println("-------------------------")
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
	} else {
		fmt.Println("Must use 'test' or 'run' as parameters.")
		utils.Info.Println("Not a proper parameter")
		os.Exit(1)
	}

}
