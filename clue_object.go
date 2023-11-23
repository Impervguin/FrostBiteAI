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
		fmt.Println(string(row))
	}
}

func (clue ClueObject) action() {
	content, err := os.ReadFile(clue.filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	var data clueData

	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return
	}
	StartOneSideInteface(data)
}
