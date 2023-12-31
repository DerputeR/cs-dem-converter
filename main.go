package main

import (
	"bufio"
	"demo-parser/dem"
	"demo-parser/utils"
	"fmt"
	"os"
	"strings"
	//"strconv"
)

var demoPath string = ""
var scanner *bufio.Scanner

func set_demofile(scanner *bufio.Scanner) bool {
	fmt.Println("---")
	fmt.Println("Enter demo path (enter empty string to exit).")

	var validInput bool = false
	for !validInput {
		scanner.Scan()
		demoPath = scanner.Text()
		demoPath = strings.TrimSpace(demoPath)

		if demoPath == "" {
			return false
		}

		if len(demoPath) > 2 && demoPath[0] == '"' && demoPath[len(demoPath)-1] == '"' {
			demoPath = demoPath[1 : len(demoPath)-1]
		}

		validInput = utils.CheckFile(demoPath)
		if !validInput {
			fmt.Print("File not found at ", demoPath, ", please try again.\n")
		}
	}
	return true
}

func setup_parser(scanner *bufio.Scanner) bool {
	if set_demofile(scanner) {
		fmt.Println()
		return true
	}
	return false
}

func main() {
	// create scanner and save it
	scanner = bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Demo Parser.")

	if setup_parser(scanner) {
		fmt.Print("Opening ", demoPath, "\n")
		dem.LoadDem(demoPath)
	}
}
