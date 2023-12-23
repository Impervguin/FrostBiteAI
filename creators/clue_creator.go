package main

// go run main.go <путь к файлу>

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// clueData - структура для хранения данных из файла
type clueData struct {
	Startline string   // Первая строка файла
	ASCII_mtr [][]byte // Матрица ASCII-арт из остальных строк файла
}

func main() {
	// Проверка, что передан единственный аргумент (путь к файлу)
	if len(os.Args) != 2 {
		fmt.Println("Необходимо передать путь к файлу в качестве аргумента.")
		return
	}

	// Получение пути к файлу из аргументов командной строки
	filePath := os.Args[1]

	// Чтение содержимого файла
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Разделение содержимого файла на строки
	lines := strings.Split(string(content), "\n")

	// Проверка наличия данных в файле
	if len(lines) < 1 {
		fmt.Println("Файл пуст.")
		return
	}

	// Инициализация структуры clueData
	data := clueData{
		Startline: lines[0],
	}

	// Заполнение матрицы ASCII-арт из остальных строк файла
	for i := 1; i < len(lines); i++ {
		data.ASCII_mtr = append(data.ASCII_mtr, []byte(lines[i]))
	}

	// Преобразование данных в формат JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Ошибка преобразования в JSON:", err)
		return
	}

	// Формирование пути к JSON-файлу
	jsonFilePath := strings.TrimSuffix(filePath, ".txt") + ".json"

	// Запись данных в файл JSON
	err = ioutil.WriteFile(jsonFilePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Программа успешно завершена. Результат записан в", jsonFilePath)
}
