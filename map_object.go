package main

type mapObject interface {
	action()
}

type mapItem struct {
	char byte
	color ColorCode
	objName string
	passable bool
}

var Wall = mapItem{'#', ColorRed, "Wall", false}
var Empty = mapItem{' ', ColorReset, "Empty", true}
var Path = mapItem{'.', ColorReset, "Path", true}
var Door = mapItem{'0', ColorYellow, "Door", false}
var Clue = mapItem{'?', ColorReset, "Clue", false}

func GetMapItem(char byte) mapItem {
	switch char {
	case '#':
		return Wall
	case ' ':
		return Empty
	case '.':
		return Path
	case '0':
		return Door
	case '?':
		return Clue
	default:
		return mapItem{char, ColorReset, "Item", false}
	}
}

func PrintMapItem(item mapItem) {
	PrintColorChar(item.char, item.color)
}