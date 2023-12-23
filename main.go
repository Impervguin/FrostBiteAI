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
	InitFinalObject(&m.objs)

	m.print_height, m.print_width = 20, 65
	m.player.x = 32
	m.player.y = 20
	GameStart()
	var ans int = 0
	for ans == 0 {
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
