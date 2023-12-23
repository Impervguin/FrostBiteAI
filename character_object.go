package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// characterObject - структура, представляющая объект персонажа в игре
type characterObject struct {
	CharacterID int                  // Идентификатор персонажа
	X, Y        int                  // Координаты персонажа на игровой карте
	SettingFile string               // Имя файла с настройками персонажа
	Messages    *[]map[string]string // Указатель на слайс сообщений для взаимодействия с игроком
}

// characterData - структура для хранения данных о персонаже
type characterData struct {
	ASCII_mtr   [][]byte             // Матрица ASCII-арт персонажа
	Name        string               // Имя персонажа
	Prof        string               // Профессия персонажа
	TimeStayed  string               // Время пребывания персонажа в деревне
	Info        string               // Информация о персонаже
	InitMessage string               // Начальное сообщение в диалоге персонажа
	EndMessage  string               // Сообщение при окончании диалога
	ErrMessage  string               // Сообщение при ошибке генерации
	IsKiller    bool                 // Является ли персонаж убийцей
	Answers     map[string]string    // Ответы персонажа на вопросы игрока
	Is_end      bool                 // Флаг окончания диалога
	Messages    *[]map[string]string // Указатель на слайс сообщений для взаимодействия с игроком
}

// get_answer - метод для получения ответа персонажа на вопрос игрока
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
	if err != nil {
		return name_prefix + char.ErrMessage
	}

	res, err := Get_gpt_message(char.Messages)
	if err != nil {
		return name_prefix + char.ErrMessage
	}

	return name_prefix + res
}

// get_init_message - метод для получения начального сообщения персонажа
func (char *characterData) get_init_message() string {

	var picture string = ""
	for _, line := range char.ASCII_mtr {
		picture += string(line) + "\n"
	}

	command := "[" + char.Name + "]"
	Add_user_message(char.Messages, command)
	err := Send_gpt_message(char.Messages)
	if err != nil {
		return picture + "\n" + char.InitMessage
	}
	res, err := Get_gpt_message(char.Messages)
	if err != nil {
		return picture + "\n" + char.InitMessage
	}

	return picture + "\n" + res
}

// get_end_message - метод для получения сообщения при окончании диалога персонажа
func (char *characterData) get_end_message() string {
	return char.EndMessage
}

// is_end - метод для проверки окончания диалога с персонажем
func (char *characterData) is_end() bool {
	return char.Is_end
}

// action - метод для выполнения действия персонажа
func (character characterObject) action() int {
	data := characterData{}
	content, err := os.ReadFile(character.SettingFile)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return 0
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return 0
	}
	data.Is_end = false
	data.Messages = character.Messages

	StartTwoSideInteface(&data)
	return 0
}

// get_x - метод для получения координаты X персонажа
func (character characterObject) get_x() int {
	return character.X
}

// get_y - метод для получения координаты Y персонажа
func (character characterObject) get_y() int {
	return character.Y
}

// Characters_init_message - функция для инициализации сообщений персонажей
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
