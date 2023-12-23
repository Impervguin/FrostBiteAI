package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// FinalCharacterObject представляет информацию о персонаже в финальном объекте.
type FinalCharacterObject struct {
	CharacterID int
	IsKiller    bool
	Name        string
	Prof        string
}

// FinalObject представляет финальный объект с координатами X и Y, а также картой персонажей.
type FinalObject struct {
	X, Y  int
	Chars map[int]FinalCharacterObject
}

// get_x возвращает значение X для FinalObject.
func (obj FinalObject) get_x() int {
	return obj.X
}

// get_y возвращает значение Y для FinalObject.
func (obj FinalObject) get_y() int {
	return obj.Y
}

// action выполняет действие, предлагая детективу выбрать персонажа из карты.
func (obj FinalObject) action() int {
	fmt.Println("Готов ли ты сделать выбор, детектив?")
	fmt.Println()
	for k, v := range obj.Chars {
		fmt.Println(k, " - ", v.Name, ", ", v.Prof)
	}
	var id int
	fmt.Scan(&id)
	if _, ok := obj.Chars[id]; ok {
		return id
	}
	return 0
}

// InitFinalObject инициализирует объект FinalObject на основе массива mapObject.
func InitFinalObject(objs *[]mapObject) error {
	var final *FinalObject
	Chars := map[int]FinalCharacterObject{}

	// Перебираем все объекты в массиве mapObject.
	for _, v := range *objs {
		// Проверяем, является ли текущий объект FinalObject.
		if fin, ok := v.(*FinalObject); ok {
			final = fin
		}
		// Проверяем, является ли текущий объект characterObject.
		if char, ok := v.(*characterObject); ok {
			// Проверяем, есть ли уже информация о персонаже с таким CharacterID.
			if _, ok := Chars[char.CharacterID]; ok {
				continue
			}

			// Читаем данные о персонаже из файла и преобразуем их из JSON.
			data := characterData{}
			content, err := os.ReadFile(char.SettingFile)
			if err != nil {
				return fmt.Errorf("Ошибка чтения файла")
			}
			err = json.Unmarshal(content, &data)
			if err != nil {
				return fmt.Errorf("Ошибка разбора JSON")
			}

			// Добавляем информацию о персонаже в карту Chars.
			Chars[char.CharacterID] = FinalCharacterObject{CharacterID: char.CharacterID, Name: data.Name, Prof: data.Prof, IsKiller: data.IsKiller}
		}
	}

	// Проверяем, был ли найден объект FinalObject.
	if final == nil {
		return fmt.Errorf("Ошибка: нет объекта финала")
	}

	// Устанавливаем карту персонажей в объекте FinalObject.
	final.Chars = Chars
	return nil
}
