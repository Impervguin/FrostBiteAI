package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type characterData struct {
	ASCII_mtr   [][]byte
	Name        string
	Prof        string
	TimeStayed  string
	Info        string
	InitMessage string
	EndMessage  string
	ErrMessage  string
	Answers     map[string]string
}

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите название записываемого файла: ")
	Scanner.Scan()
	jsonname := Scanner.Text()
	
	JsonFile, err := os.Create(jsonname)
	if (err != nil) {
		fmt.Println("Ошибка файла:", err)
		return
	}

	fmt.Print("Введите имя персонажа: ")
	Scanner.Scan()
	name := Scanner.Text()


	fmt.Print("Введите имя файла с картинкой персонажа: ")
	Scanner.Scan()
	fname := Scanner.Text()

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
	for i := 1; i < len(lines); i++ {
		ASCII_mtr = append(ASCII_mtr, []byte(lines[i]))
	}

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

	fmt.Println("Далее введите заготовленные реплики на вопросы персонажа. (Пустая строка - окончание реплик)")
	m := make(map[string]string)
	for {

		fmt.Print("Введите реплику игрока: ")
		Scanner.Scan()
		key := Scanner.Text()
		if (len(key) == 0) {
			break
		}
		fmt.Print("Введите ответ персонажа: ")
		Scanner.Scan()
		val := Scanner.Text()
		if (len(val) == 0) {
			break
		}
		m[key] = val
	}


	data := characterData{ASCII_mtr:ASCII_mtr, Name:  name, Prof: prof, TimeStayed:  time_stayed, Info: info, InitMessage: init,EndMessage: end, Answers: m, ErrMessage: errmes}
	
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Ошибка преобразования в JSON:", err)
		return
	}
	
	_, err = JsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
}