package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type mapObject interface {
	action() int
	get_x() int
	get_y() int
}

type TwoSideInterface interface {
	get_answer(string) string
	get_init_message() string
	get_end_message() string
	is_end() bool
}

func StartTwoSideInteface(inter TwoSideInterface) {
	fmt.Println(inter.get_init_message())
	scan := bufio.NewScanner(os.Stdin)
	nend := true
	for nend {
		fmt.Print("Игрок: ")
		scan.Scan()
		player_mes := scan.Text()
		ans := inter.get_answer(player_mes)
		fmt.Println(ans)
		nend = inter.is_end()
	}
	fmt.Print(inter.get_end_message())
	fmt.Scanln()
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type OneSideInterface interface {
	print_message()
	print_picture()
}

func StartOneSideInteface(inter OneSideInterface) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	inter.print_picture()
	inter.print_message()
	fmt.Scanln()
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type mapItem struct {
	char     byte
	color    ColorCode
	objName  string
	passable bool
}

var WallItem = mapItem{'#', ColorRed, "Wall", false}
var ChristItem = mapItem{'$', ColorYellow, "Wall", false}
var EmptyItem = mapItem{' ', ColorReset, "Empty", true}
var PathItem = mapItem{'.', ColorReset, "Path", true}
var DoorItem = mapItem{'0', ColorYellow, "Door", false}
var ClueItem = mapItem{'?', ColorReset, "Clue", false}

func GetMapItem(char byte) mapItem {
	switch char {
	case '#':
		return WallItem
	case '$':
		return ChristItem
	case ' ':
		return EmptyItem
	case '.':
		return PathItem
	case '0':
		return DoorItem
	case '?':
		return ClueItem
	default:
		return mapItem{char, ColorReset, "Item", false}
	}
}

func PrintMapItem(item mapItem) {
	PrintColorChar(item.char, item.color)
}
