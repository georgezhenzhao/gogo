package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
)

// CheckPidInfo
func CheckPidInfo(p string) {

	fmt.Printf("processID: %s \n ", p)

	pid, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	process, err := os.FindProcess(int(pid))
	if err != nil {
		fmt.Printf("Failed to find process: %s\n", err)
	} else {
		err := process.Signal(syscall.Signal(0))
		fmt.Printf("process.Signal on pid %d returned: %v\n", pid, err)
	}
}
