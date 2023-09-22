package utils

import (
	"os"
)

type Game int8

const (
	GameNone Game = iota
	GameCSGO
	GameCS2
)

func (game Game) String() string {
	switch game {
	case GameNone:
		return "none"
	case GameCSGO:
		return "csgo"
	case GameCS2:
		return "cs2"
	}
	return "unknown"
}

func CheckFile(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	f.Close()
	return true
}
