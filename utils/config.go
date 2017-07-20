package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Users       []string
	Loggingfile string
}

// Initilize log
func getLogFile() string {

	file, _ := os.Open("conf/conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	errconf := decoder.Decode(&configuration)
	if errconf != nil {
		fmt.Println("errconf: ", errconf)
	}
	fmt.Println(configuration.Loggingfile)
	return configuration.Loggingfile
}
