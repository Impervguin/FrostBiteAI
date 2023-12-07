package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var WIN_TEXT string = `
                                   .''.       
       .''.      .        *''*    :_\/_:     . 
      :_\/_:   _\(/_  .:.*_\/_*   : /\ :  .'.:.'.
  .''.: /\ :   ./)\   ':'* /\ * :  '..'.  -=:o:=-
 :_\/_:'.:::.    ' *''*    * '.\'/.' _\(/_'.':'.'
 : /\ : :::::     *_\/_*     -= o =-  /)\    '  *
  '..'  ':::'     * /\ *     .'/.\'.   '
      *            *..*         :
        *
        *
`

var LOSER_TEXT string = `
⠀⠀⠀⠀⠀⠀⢀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠂⡰⡆⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠠⣞⣟⣹⢧⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⣿⡾⢛⢻⣷⠀⠀⠀⠀⠀
⠀⠀⠀⢀⣾⡟⠛⢻⣷⣝⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣤⣤⣤⣀⣀⣀⣾⣧⡿⢉⠰⢠⢨⣿⠁⠀⠀⠀⠀
⠀⠠⣯⣹⣿⢀⠣⠠⡙⣿⣌⢷⣦⣤⣀⣤⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠠⣾⣿⣿⠿⣿⣿⣹⣿⣿⡟⡁⠆⡑⢂⡼⣿⣠⡀⠀⠀⠀
⠀⢐⣿⡿⢟⣆⠢⠑⡄⢹⣿⣿⣿⣿⠛⠍⡛⣿⣶⣤⠂⠀⠀⠀⠀⠀⠀⠀⣯⣿⣇⠂⠌⢿⣿⢿⡟⠡⠐⡌⡐⢢⠳⢉⡉⢿⣦⠀⠀
⣿⣿⡃⡐⡈⠻⣧⢈⠰⢀⠛⣿⣿⠃⠌⡒⣱⣿⣷⡏⠀⠀⠀⠀⠀⠀⠀⠀⠻⣞⣿⣎⠰⢈⠙⡋⠔⣁⠣⡐⠡⡉⠔⠢⡐⢈⣿⠆⠀
⣷⣿⣡⠐⠤⡑⡀⠆⡑⠢⠌⣈⢉⡐⢂⠱⣼⣏⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢾⣿⢷⡄⠊⡔⠡⢂⠔⠡⡁⠆⡉⢆⠁⣞⢻⣦⡀
⣿⣾⣿⣶⠆⡑⠌⢂⠡⢃⠒⢄⠢⠘⢦⣷⣿⡞⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡻⣾⣿⣔⢀⠃⢆⡘⠰⠠⠑⡰⢀⠎⣀⠂⢸⣿
⣿⣿⡋⠄⢢⠑⡌⠂⡅⢊⠔⡨⢀⢃⣾⡟⡛⠷⡴⣋⣭⣷⣶⣶⣶⣶⣶⣦⣶⣾⣭⣽⣚⢾⡿⣷⣬⡀⠌⣁⠣⠑⡐⠢⢈⠤⠑⣼⡿
⣿⡹⣷⣌⠄⠃⡌⠰⡈⠆⢂⠡⣢⣾⣿⢃⣷⡿⠟⢛⠩⢉⣷⠔⡠⢶⣄⡐⡈⠌⣉⠙⢛⠻⢷⣮⣟⠻⢷⣶⣤⣧⣤⣑⣦⣤⣿⠟⠁
⠀⠉⠻⣽⣻⢷⣄⣃⣐⣨⣴⣿⠟⣃⣴⡟⠩⢐⠨⣤⣾⡿⢋⠰⠀⢭⠻⢿⣶⣬⣤⣘⣠⣈⠄⡉⠻⢿⣮⠻⡿⠿⠉⠉⠛⠋⠁⠀⠀
⠀⠀⠀⠀⠈⠙⠽⠿⠛⢻⣭⣷⣾⡟⠳⠶⢷⠾⠟⢋⣱⣶⢆⢣⡙⢾⣷⣤⡈⢉⠍⠛⠛⢋⡐⠌⡐⢂⠹⣿⣌⢤⡀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⣿⡿⢃⠐⡢⢁⠆⠰⣨⣿⣿⢏⡜⣢⢝⣚⣿⣿⣿⣤⠊⠔⡑⠢⠌⢢⠁⢆⠡⢈⢿⣦⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣼⢳⡿⠁⢆⠱⢀⢃⣬⣷⣿⡿⢯⡟⡼⠥⡏⢞⠳⣟⣿⣿⣷⣮⣤⣁⠊⡄⠡⢊⠐⡂⠌⣿⡆⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠰⣷⣿⠃⡜⠠⣳⣾⣿⣿⢿⡻⡵⢏⠞⡴⢫⢜⡣⡝⡸⢧⣟⣻⠿⣿⣿⣿⣿⣇⠂⡱⢈⠰⢸⣿⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣼⣻⡟⢰⠠⠉⣿⡯⡝⣜⡲⣿⠈⠎⡽⢌⡳⢌⡖⡍⠖⡃⠌⣹⡿⣔⡲⡱⢾⣧⠂⢡⠊⡄⢹⣿⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣿⡇⢪⠇⡈⣿⣱⢹⢦⡹⣿⠀⢍⣰⣦⣶⣷⣶⣾⣴⣀⠃⠼⣿⡰⢳⣍⢻⣷⠈⢢⠑⡠⢹⣿⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣿⠥⡙⢆⠡⣿⣥⢫⡖⣭⣿⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡈⣿⡱⣏⠼⣹⣯⠐⢂⠱⣀⢹⣿⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣿⡇⡘⡅⢂⣿⢖⡳⣚⢴⣻⣿⣿⣿⣻⢏⡿⣽⣻⣿⣿⣿⣿⡼⢳⢭⡚⡽⣷⠈⢌⡒⢤⣿⠇⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢻⣧⠐⡌⠄⣿⢎⡵⣍⢮⣽⣿⣿⣷⠿⠿⠿⢷⣷⣿⣿⣿⣿⡗⡹⢶⡹⣜⣿⠀⢊⢥⣾⠏⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⣇⠌⢂⣿⣚⠼⡜⢦⣿⠟⡉⠄⣒⠨⣐⠄⡠⠌⡙⠛⢻⣭⡓⣧⠳⣼⡏⢰⣡⡿⣋⠄⠀⠀⠀⠀⠀⠀⠀
⠂⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣷⣄⢺⣥⢻⡙⣎⣿⠀⡔⠡⣈⡇⠘⣆⠐⢢⠐⡉⢸⣷⡹⣜⠳⣜⣧⣿⢋⠠⠃⠀⠀⠀⠀⠀⠀⠀⠀
⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣻⢷⣮⣧⡙⡖⣿⠐⠨⠑⡠⠉⡔⠠⡉⠄⠃⢌⠘⣷⣳⣼⡿⢿⡋⠈⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⡁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠓⠯⠛⠿⢷⣧⣬⣥⣆⣡⣒⣤⣁⣦⣭⣼⣶⠿⣟⣏⠷⠞⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣈⣙⢸⣉⡍⠛⠙⠛⠛⢻⡙⡏⢡⢒⣼⣉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
`



func GameFinal(id int, objs *[]mapObject) error {
	// fmt.Println("Вфывц")
	var char characterObject
	for _, v := range *objs {
		// fmt.Println(v)
		if c, ok := v.(*characterObject); ok {
			if c.CharacterID == id {
				char = *c
				break
			}
		}
	}

	data := characterData{}
	content, err := os.ReadFile(char.SettingFile)
	if err != nil {
		return fmt.Errorf("Ошибка чтения файла")
	}
	// fmt.Println(err)
	err = json.Unmarshal(content, &data)
	// fmt.Println(err)
	if err != nil {
		return fmt.Errorf("Ошибка разбора JSON")
	}

	// fmt.Println(data)
	if data.IsKiller {
		fmt.Println(WIN_TEXT)
		fmt.Println("Вы Разгадали это дело!", data.Name, "- убийца!")
	} else {
		fmt.Println(LOSER_TEXT)
		fmt.Println("Вы повесили невинного:", data.Name, ", убийства в Лунном зеркальце продолжатся.")
	}
	return nil
}