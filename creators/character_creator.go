package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type characterData struct {
	ASCII_mtr    [][]byte
	Name         string
	Prof         string
	TimeStayed  string
	Info         string
	InitMessage string
	Answers      map[string]string
}

func main() {
	fmt.Print("Введите название записываемого файла: ")
	var jsonname string
	fmt.Scanln(&jsonname)
	
	JsonFile, err := os.Create(jsonname)
	if (err != nil) {
		fmt.Println("Ошибка файла:", err)
		return
	}

	fmt.Print("Введите имя персонажа: ")
	var name string
	fmt.Scanln(&name)


	fmt.Print("Введите имя файла с картинкой персонажа: ")
	var fname string
	fmt.Scanln(&fname)

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
	var prof string
	fmt.Scanln(&prof)

	fmt.Print("Введите время нахождения персонажа в деревне: ")
	var time_stayed string
	fmt.Scanln(&time_stayed)

	fmt.Print("Введите начальное сообщение в диалоге персонажа: ")
	var init string
	fmt.Scanln(&init)

	fmt.Print("Введите информацию о персонаже: ")
	var info string
	fmt.Scanln(&info)

	fmt.Println("Далее введите заготовленные реплики на вопросы персонажа. (Пустая строка - окончание реплик)")
	m := make(map[string]string)
	for {
		var key, val string

		fmt.Print("Введите реплику игрока: ")
		fmt.Scanln(&key)
		if (len(key) == 0) {
			break
		}
		fmt.Print("Введите ответ персонажа: ")
		fmt.Scanln(&val)
		if (len(val) == 0) {
			break
		}
		m[key] = val
	}


	data := characterData{ASCII_mtr:ASCII_mtr, Name:  name, Prof: prof, TimeStayed:  time_stayed, Info: info, InitMessage: init, Answers: m}
	
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