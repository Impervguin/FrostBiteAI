package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ClueObject struct {
	x, y     int
	filename string
}

type clueData struct {
	Startline string
	ASCII_mtr [][]byte
}

func (clue clueData) print_message() {
	fmt.Println(clue.Startline)
}

func (clue clueData) print_picture() {
	for _, row := range clue.ASCII_mtr {
		for _, char := range row {
			if char == '#' {
				PrintColorChar(char, ColorRed)
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

func (clue ClueObject) action() int {
	content, err := os.ReadFile(clue.filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return 0
	}

	var data clueData

	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return 0
	}
	StartOneSideInteface(data)
	return 0
}

func (clue ClueObject) get_x() int {
	return clue.x
}

func (clue ClueObject) get_y() int {
	return clue.y
}
