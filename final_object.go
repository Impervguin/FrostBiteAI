package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FinalCharacterObject struct {
	CharacterID int
	IsKiller    bool
	Name        string
	Prof        string
}

type FinalObject struct {
	X, Y  int
	Chars map[int]FinalCharacterObject
}

func (obj FinalObject) get_x() int {
	return obj.X
}

func (obj FinalObject) get_y() int {
	return obj.Y
}

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


func InitFinalObject(objs *[]mapObject) error {
	var final *FinalObject
	Chars := map[int]FinalCharacterObject{}
	for _, v := range *objs {
		if fin, ok := v.(*FinalObject); ok {
			final = fin
		}
		if char, ok := v.(*characterObject); ok {
			if _, ok := Chars[char.CharacterID]; ok {
				continue
			}
			data := characterData{}
			content, err := os.ReadFile(char.SettingFile)
			if err != nil {
				return fmt.Errorf("Ошибка чтения файла")
			}
			err = json.Unmarshal(content, &data)
			if err != nil {
				return fmt.Errorf("Ошибка разбора JSON")
			}
			Chars[char.CharacterID] = FinalCharacterObject{CharacterID: char.CharacterID, Name: data.Name, Prof: data.Prof, IsKiller: data.IsKiller}
		}
	}
	if (final == nil) {
		return fmt.Errorf("Ошибка: нет объекта финала")
	} 
	final.Chars = Chars
	return nil
}