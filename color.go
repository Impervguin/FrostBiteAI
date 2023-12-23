package main

import "fmt"

type ColorCode string

var ColorReset ColorCode = "\033[0m"
var ColorRed ColorCode = "\033[31m"
var ColorGreen ColorCode = "\033[32m"
var ColorYellow ColorCode = "\033[33m"
var ColorBlue ColorCode = "\033[34m"
var ColorPurple ColorCode = "\033[35m"
var ColorCyan ColorCode = "\033[36m"
var ColorGray ColorCode = "\033[37m"
var ColorWhite ColorCode = "\033[97m"

// Функция PrintColorChar принимает байт символа и код цвета, затем выводит символ с применением указанного цвета.
func PrintColorChar(ch byte, color ColorCode) {
	fmt.Printf("%s%c%s", color, ch, ColorReset)
}
