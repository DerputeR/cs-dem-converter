package main

import (
	"bufio"
	"demo-parser/utils"
	"fmt"
	"os"
	"strings"
	//"strconv"
)

var parserMode utils.Game = utils.GameNone
var demoPath string = ""
var scanner *bufio.Scanner

func set_parser(scanner *bufio.Scanner) bool {
	fmt.Println("---")
	fmt.Println("1. Load CS:[g]O demo")
	fmt.Println("2. Load CS[2] demo")
	fmt.Println("3. [q]uit")

	var validInput bool = false
	var gameSelect string
	for !validInput {
		scanner.Scan()
		gameSelect = scanner.Text()
		gameSelect = strings.ToLower(gameSelect)

		switch gameSelect {
		case "1":
			validInput = true
			parserMode = utils.GameCSGO
		case "g":
			validInput = true
			parserMode = utils.GameCSGO
		case "csgo":
			validInput = true
			parserMode = utils.GameCSGO
		case "2":
			validInput = true
			parserMode = utils.GameCS2
		case "cs2":
			validInput = true
			parserMode = utils.GameCS2
		case "3":
			validInput = true
			parserMode = utils.GameNone
			return false
		case "q":
			validInput = true
			parserMode = utils.GameNone
			return false
		case "quit":
			validInput = true
			parserMode = utils.GameNone
			return false
		default:
			validInput = false
			parserMode = utils.GameNone
		}
		if !validInput {
			fmt.Print("Invalid input, please try again: ")
		}
	}
	return true
}

func set_demofile(scanner *bufio.Scanner) bool {
	fmt.Println("---")
	fmt.Println("Enter demo path (enter empty string to go back).")

	var validInput bool = false
	for !validInput {
		scanner.Scan()
		demoPath = scanner.Text()
		demoPath = strings.TrimSpace(demoPath)
		if demoPath == "" {
			// fmt.Println("entered: ", demoPath)
			return false
		}
		validInput = utils.CheckFile(demoPath)
		if !validInput {
			fmt.Print("File not found at ", demoPath, ", please try again.\n")
		}
	}
	return true
}

func setup_parser(scanner *bufio.Scanner) bool {
	for {
		if !set_parser(scanner) {
			return false
		}
		if set_demofile(scanner) {
			fmt.Println()
			break
		}
	}
	return true
}

func main() {
	// create scanner and save it
	scanner = bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Demo Parser.")

	if setup_parser(scanner) {
		fmt.Print("Parsing", demoPath, "with", parserMode, "parser...")
	}
}
