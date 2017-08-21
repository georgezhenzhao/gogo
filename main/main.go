package main

import (
	"bufio"
	"fmt"
	"os"

	"../try"
	"../utils"
	"../web"
)

func main() {

	fmt.Printf("pid: %d\n", os.Getpid())

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

		parMap := make(map[string]string)
		parMap["p1"] = os.Args[1]
		fmt.Println("parMap:", parMap)

	} else {

		fmt.Println("gogo is starting without params given.")
		fmt.Println("Nothing we can do now.")
		utils.Info.Println("Nothing we can do now.")
		os.Exit(0)
	}
	fmt.Println("********************")

	if os.Args[1] == "test" {
		try.TryFunc()
		try.TryVar()
		try.TryCont()
		try.TryPointer()

		fmt.Println(try.GV)
		fmt.Println(try.GC)

		//try.TryMethod()
		try.TryInterface()
		try.TryErrors()
		try.TryChannels()
		try.TryWorkers()
		try.TryJson()

		os.Exit(0)

	} else if os.Args[1] == "get" {

		urls := make([]string, 3)
		urls[0] = "http://172.20.8.108:9090/admin/info"
		urls = append(urls, "http://google.com", "http://localhost:9090/admin/health")
		fmt.Println("emp:", urls)

		var rs string = web.CallGet(urls[0])

		utils.Info.Println(rs)

	} else if os.Args[1] == "ps" && len(os.Args) == 3 {

		utils.CheckPidInfo(os.Args[2])

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
