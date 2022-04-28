package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkLogs(file string) bool {
	logFile, err := os.Open(file)

	if err != nil {
		fmt.Println("Could not find log file")
		return false
	}
	logFile.Close()

	return true
}

func readLogs(file string) string {
	found := checkLogs(file)
	if found != true {
		fmt.Println("Could not find file")
	}

	logFile, err := os.Open(file)
	defer logFile.Close()
	if err != nil {
		fmt.Println("Could not open File")
	}
	fileScanner := bufio.NewScanner(logFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	return "DONE"
}
