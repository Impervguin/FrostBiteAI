package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// ClueObject представляет объект-подсказку с координатами и именем файла.
type ClueObject struct {
	x, y     int
	filename string
}

// clueData представляет структуру данных, используемую для разбора JSON-файла.
type clueData struct {
	Startline string
	ASCII_mtr [][]byte
}

// print_message выводит стартовую строку подсказки.
func (clue clueData) print_message() {
	fmt.Println(clue.Startline)
}

// print_picture выводит ASCII-арт из матрицы в подсказке с поддержкой цветов.
func (clue clueData) print_picture() {
	for _, row := range clue.ASCII_mtr {
		for _, char := range row {
			if char == '#' {
				PrintColorChar(char, ColorRed) // Предположим, что PrintColorChar и ColorRed определены в другом месте в коде.
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

// action выполняет действие, считывая данные из файла и вызывая StartOneSideInteface.
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
	StartOneSideInteface(data) // Предположим, что StartOneSideInteface определена в другом месте в коде.
	return 0
}

// get_x возвращает координату X подсказки.
func (clue ClueObject) get_x() int {
	return clue.x
}

// get_y возвращает координату Y подсказки.
func (clue ClueObject) get_y() int {
	return clue.y
}
