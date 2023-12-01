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
	Messages    *[]map[string]string
}

type characterData struct {
	ASCII_mtr   [][]byte
	Name        string
	Prof        string
	TimeStayed  string
	Info        string
	InitMessage string
	EndMessage  string
	Answers     map[string]string
	Is_end      bool
}

func (char *characterData) get_answer(question string) string {

	ans, ok := char.Answers[question]
	if question == "q" {
		char.Is_end = true
		return ""
	}

	if ok {
		return ans
	}
	return "Я не знаю что на это ответить."
}

func (char *characterData) get_init_message() string {
	res := ""
	for _, str := range char.ASCII_mtr {
		res += string(str) + "\n"
	}
	res += char.InitMessage
	return res
}

func (char *characterData) get_end_message() string {
	return char.EndMessage
}

func (char *characterData) is_end() bool {
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
	StartTwoSideInteface(&data)
}

func (character characterObject) get_x() int {
	return character.X
}

func (character characterObject) get_y() int {
	return character.Y
}

func Characters_init_message(objs []mapObject) *[]map[string]string {
	res_str := "Твои роли:\n"
	data := characterData{}
	res := []map[string]string{}
	for _, obj := range objs {
		if c_obj, ok := obj.(*characterObject); ok {
			c_obj.Messages = &res
			content, err := os.ReadFile(c_obj.SettingFile)
			if err != nil {
				fmt.Println("Ошибка чтения файла:", err)
				return nil
			}
			err = json.Unmarshal(content, &data)
			if err != nil {
				fmt.Println("Ошибка разбора JSON:", err)
				return nil
			}
			res_str += fmt.Sprintf("%s - %s - Время пребывания в деревне: %s - %s\n", data.Name, data.Prof, data.TimeStayed, data.Info)
		}
	}
	res_str += "Твоя цель вести диалог со мной отыгрывая одного из этих персонажей, которого я буду выбирать. Когда я буду писать имя персонажа в квадратных скобках начинай его отыгрывать.\n"
	res = append(res, map[string]string{"role": "user", "content": res_str})
	return &res
}
