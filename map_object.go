package main

import "fmt"

type mapObject interface {
	action()
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
	for !inter.is_end() {
		fmt.Print("Игрок: ")
		var player_mes string
		fmt.Scanln((&player_mes))
		ans := inter.get_answer(player_mes)
		fmt.Println(ans)
	}
	fmt.Println(inter.get_end_message())
	fmt.Scanln()
}

type OneSideInterface interface {
	print_message()
	print_picture()
}

func StartOneSideInteface(inter OneSideInterface) {
	inter.print_message()
	inter.print_picture()
	fmt.Scanln()
}

type CharacterObject struct {
	characterid int
	x, y        int
	settingfile string
}

type FinalObject struct {
	x, y        int
	settingfile string
}

func (character CharacterObject) action() {
	fmt.Println("Characters are under construction, sorry!")
}

func (character CharacterObject) get_x() int {
	return character.x
}

func (character CharacterObject) get_y() int {
	return character.y
}

func (final FinalObject) action() {
	fmt.Println("Final are under construction, sorry!")
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
