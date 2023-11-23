package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type characterObject struct {
	CharacterID int
	X, Y        int
	SettingFile string
}

type characterData struct {
	ASCII_mtr   [][]byte
	Name        string
	Prof        string
	TimeStayed  string
	Info        string
	InitMessage string
	Answers     map[string]string
	Is_end      bool
}

func (char characterData) get_answer(question string) string {
	ans, ok := char.Answers[question]
	if ok && question == "q" {
		char.Is_end = true
	}
	if ok {
		return ans
	}
	return "Я не знаю что на это ответить."
}

func (char characterData) get_init_message() string {
	res := ""
	for _, str := range char.ASCII_mtr {
		res += string(str) + "\n"
	}
	res += char.InitMessage
	return res
}

func (char characterData) get_end_message() string {
	return ""
}

func (char characterData) is_end() bool {
	return char.Is_end
}

func (character characterObject) action() {
	data := characterData{}
	content, err := os.ReadFile(character.SettingFile)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return
	}
	data.Is_end = false
	StartTwoSideInteface(data)
}

func (character characterObject) get_x() int {
	return character.X
}

func (character characterObject) get_y() int {
	return character.Y
}

