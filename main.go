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
	m.objs = items
	// 	fmt.Println(items)
	InitFinalObject(&m.objs)
	// fmt.Println(Send_gpt_message(res))
	// fmt.Println(Get_gpt_message(res))
	// *res = append(*res, map[string]string{"role": "user", "content" : "Как прошел твой день, Максимилью?"})
	// fmt.Println(Send_gpt_message(res))
	// fmt.Println(Get_gpt_message(res))

	m.print_height, m.print_width = 20, 65
	m.player.x = 32
	m.player.y = 20
	// fmt.Println(m.objs)
	GameStart()
	var ans int = 0
	for ans == 0 {
		// for _, v := range m.objs {
		// 	if c, ok := v.(*characterObject); ok {
		// 		fmt.Println(*c)
		// 	}

		// }
		// fmt.Println(m.objs)
		// fmt.Println(m.player)
		m.print_map()
		tty, _ := tty.Open()
		r, _ := tty.ReadRune()
		tty.Close()
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		if r == 'd' {
			ans, _ = m.PlayerMoveRight()
		}
		if r == 's' {
			ans, _ = m.PlayerMoveDown()
		}
		if r == 'w' {
			ans, _ = m.PlayerMoveUp()
		}
		if r == 'a' {
			ans, _ = m.PlayerMoveLeft()
		}
	}
	GameFinal(ans, &m.objs)
}
