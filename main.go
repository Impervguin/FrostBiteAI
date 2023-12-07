package main

import (
	"os"
	"os/exec"

	"github.com/mattn/go-tty"
)

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	m, _ := ReadMapFromFile("./data/maps/map.txt")
	items, _ := ReadMapItems("./data/maps/objects.txt")
	Characters_init_message(items)
	// fmt.Println(Send_gpt_message(res))
	// fmt.Println(Get_gpt_message(res))
	// *res = append(*res, map[string]string{"role": "user", "content" : "Как прошел твой день, Максимилью?"})
	// fmt.Println(Send_gpt_message(res))
	// fmt.Println(Get_gpt_message(res))


	m.objs = items
	m.print_height, m.print_width = 20, 65
	m.player.x = 32
	m.player.y = 20
	// fmt.Println(m.objs)
	GameStart()
	for true {
		m.print_map()
		tty, _ := tty.Open()
		r, _ := tty.ReadRune()
		tty.Close()
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		if r == 'd' {
			m.PlayerMoveRight()
		}
		if r == 's' {
			m.PlayerMoveDown()
		}
		if r == 'w' {
			m.PlayerMoveUp()
		}
		if r == 'a' {
			m.PlayerMoveLeft()
		}

	}

}
