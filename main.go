package main

import (
	"demo-parser/utils"
	"fmt"
	"strings"
	//"strconv"
)

var parserMode utils.Game = utils.GameNone
var demoPath string = ""

func set_parser() bool {
	fmt.Println("Please make a selection.")
	fmt.Println("1. Load CS:[g]O demo")
	fmt.Println("2. Load CS[2] demo")
	fmt.Println("3. Quit")

	var validInput bool = false
	var gameSelect string
	for !validInput {
		fmt.Scan(&gameSelect)
		gameSelect = strings.ToLower(gameSelect)

		switch gameSelect {
		case "1":
			validInput = true
			parserMode = utils.GameCSGO
		case "2":
			validInput = true
			parserMode = utils.GameCS2
		case "g":
			validInput = true
			parserMode = utils.GameCSGO
		case "3":
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

func set_demofile() bool {
	fmt.Print("Demo path: (type \"cancel\" with no quotes to go back) ")

	var validInput bool = false
	for !validInput {
		fmt.Scan(&demoPath)
		if demoPath == "cancel" {
			return false
		}
		validInput = utils.CheckFile(demoPath)
		if !validInput {
			fmt.Print("File not found at ", demoPath, ", please try again: ")
		}
	}
	return true
}

func setup_parser() bool {
	for {
		if !set_parser() {
			return false
		}
		if set_demofile() {
			break
		}
	}
	return true
}

func main() {
	fmt.Println("Welcome to the Demo Parser.")

	if setup_parser() {
		fmt.Print("Parsing", demoPath, "with", parserMode, "parser...")
	}
}
