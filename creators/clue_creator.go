package main

// go run main.go <путь к файлу>

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type clueData struct {
	Startline string
	ASCII_mtr [][]byte
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	filePath := os.Args[1]

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	if len(lines) < 1 {
		fmt.Println("Файл пуст")
		return
	}

	data := clueData{
		Startline: lines[0],
	}

	for i := 1; i < len(lines); i++ {
		data.ASCII_mtr = append(data.ASCII_mtr, []byte(lines[i]))
	}

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Ошибка преобразования в JSON:", err)
		return
	}

	jsonFilePath := strings.TrimSuffix(filePath, ".txt") + ".json"
	err = ioutil.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Программа успешно завершена. Результат записан в", jsonFilePath)
}

