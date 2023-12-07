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
	ErrMessage  string
	Answers     map[string]string
	Is_end      bool
	Messages    *[]map[string]string
}

func (char *characterData) get_answer(question string) string {
	name_prefix := char.Name + ": "

	ans, ok := char.Answers[question]
	if question == "q" {
		char.Is_end = true
		return ""
	}

	if ok {
		return name_prefix + ans
	}
	
	Add_user_message(char.Messages, question)
	err := Send_gpt_message(char.Messages)
	if (err != nil) {
		return name_prefix + char.ErrMessage
	}

	res, err := Get_gpt_message(char.Messages)
	if (err != nil) {
		return name_prefix + char.ErrMessage
	}

	return name_prefix + res
}

func (char *characterData) get_init_message() string {

	var picture string = ""
	for _, line := range char.ASCII_mtr {
		picture += string(line) + "\n"
	}

	command := "[" +  char.Name  +"]"
	Add_user_message(char.Messages, command)
	err := Send_gpt_message(char.Messages)
	if (err != nil) {
		return picture + "\n" + char.InitMessage
	}
	res, err := Get_gpt_message(char.Messages)
	if (err != nil) {
		return picture + "\n" + char.InitMessage
	}

	return picture + "\n" + res
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
	data.Messages = character.Messages
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
	res_str += "Твоя цель вести диалог со мной отыгрывая одного из этих персонажей, которого я буду выбирать. Когда я буду писать имя персонажа в квадратных скобках начинай его отыгрывать. Не выводи никакой информации кроме реплик персонажей.\n"
	res = append(res, map[string]string{"role": "user", "content": res_str})
	return &res
}
