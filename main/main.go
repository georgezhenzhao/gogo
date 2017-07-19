package main

import (
	"bufio"
	"fmt"
	"os"

	"../utils"
	"../web"
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

	} else if os.Args[1] == "get" {

		urls := make([]string, 3)
		urls[0] = "http://google.com"
		urls = append(urls, "http://localhost:909/admin", "http://localhost:909/admin/health")
		fmt.Println("emp:", urls)

		var rs string = web.CallGet(urls[0])

		utils.Info.Println(rs)

	} else if os.Args[1] == "start" {

		go web.StartHttpSrv()

		fmt.Println("Please enter 'q' to exit.")
		fmt.Println("-------------------------")
		scanner := bufio.NewScanner(os.Stdin)
		var text string
		for text != "q" {
			scanner.Scan()
			text = scanner.Text()
			if text != "q" {
				fmt.Println("Your text was: ", text)
			}
		}

	} else {

		fmt.Println("Must use 'test' or 'start' as parameters.")
		utils.Info.Println("Not a proper parameter")
		os.Exit(1)
	}

}
