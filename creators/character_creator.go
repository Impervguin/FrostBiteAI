package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// characterData - структура для хранения данных о персонаже
type characterData struct {
	ASCII_mtr   [][]byte          // Матрица ASCII-арт персонажа
	Name        string            // Имя персонажа
	Prof        string            // Профессия персонажа
	TimeStayed  string            // Время пребывания персонажа в деревне
	Info        string            // Информация о персонаже
	InitMessage string            // Начальное сообщение в диалоге персонажа
	EndMessage  string            // Сообщение при окончании диалога
	ErrMessage  string            // Сообщение при ошибке генерации
	IsKiller    bool              // Является ли персонаж убийцей
	Answers     map[string]string // Ответы персонажа на вопросы игрока
}

func main() {
	Scanner := bufio.NewScanner(os.Stdin)

	// Запрос на ввод названия файла
	fmt.Print("Введите название записываемого файла: ")
	Scanner.Scan()
	jsonname := Scanner.Text()

	// Создание файла для записи JSON
	JsonFile, err := os.Create(jsonname)
	if err != nil {
		fmt.Println("Ошибка файла:", err)
		return
	}

	// Ввод основных данных о персонаже
	fmt.Print("Введите имя персонажа: ")
	Scanner.Scan()
	name := Scanner.Text()

	fmt.Print("Введите имя файла с картинкой персонажа: ")
	Scanner.Scan()
	fname := Scanner.Text()

	// Чтение файла с картинкой персонажа
	content, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	if len(lines) < 1 {
		fmt.Println("Файл пуст.")
		return
	}

	var ASCII_mtr [][]byte
	// Заполнение матрицы ASCII-арт
	for i := 1; i < len(lines); i++ {
		ASCII_mtr = append(ASCII_mtr, []byte(lines[i]))
	}

	// Ввод дополнительных данных о персонаже
	fmt.Print("Введите профессию персонажа: ")
	Scanner.Scan()
	prof := Scanner.Text()

	fmt.Print("Введите время нахождения персонажа в деревне: ")
	Scanner.Scan()
	time_stayed := Scanner.Text()

	fmt.Print("Введите начальное сообщение в диалоге персонажа: ")
	Scanner.Scan()
	init := Scanner.Text()

	fmt.Print("Введите сообщение персонажа при окончании диалога: ")
	Scanner.Scan()
	end := Scanner.Text()

	fmt.Print("Введите информацию о персонаже: ")
	Scanner.Scan()
	info := Scanner.Text()

	fmt.Print("Введите сообщение персонажа при ошибки генерации: ")
	Scanner.Scan()
	errmes := Scanner.Text()

	fmt.Print("Является ли персонаж убийцей(Да/Нет): ")
	Scanner.Scan()
	killerstr := Scanner.Text()
	killer := false
	if killerstr == "Да" {
		killer = true
	}

	// Ввод реплик и ответов персонажа на вопросы игрока
	fmt.Println("Далее введите заготовленные реплики на вопросы персонажа. (Пустая строка - окончание реплик)")
	m := make(map[string]string)
	for {
		fmt.Print("Введите реплику игрока: ")
		Scanner.Scan()
		key := Scanner.Text()
		if len(key) == 0 {
			break
		}
		fmt.Print("Введите ответ персонажа: ")
		Scanner.Scan()
		val := Scanner.Text()
		if len(val) == 0 {
			break
		}
		m[key] = val
	}

	// Создание экземпляра структуры characterData с введенными данными
	data := characterData{ASCII_mtr: ASCII_mtr, Name: name, Prof: prof, TimeStayed: time_stayed, Info: info, InitMessage: init, EndMessage: end, Answers: m, ErrMessage: errmes, IsKiller: killer}

	// Преобразование данных в формат JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Ошибка преобразования в JSON:", err)
		return
	}

	// Запись данных в файл
	_, err = JsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
}
