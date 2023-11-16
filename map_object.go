package main

import "fmt"

type mapObject interface {
	action()
	get_x() int
	get_y() int
}

type ClueObject struct {
	x, y     int
	filename string
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

func (clue ClueObject) action() {
	fmt.Println("Clue objects under construction, sorry!")
}

func (character CharacterObject) action() {
	fmt.Println("Characters are under construction, sorry!")
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
