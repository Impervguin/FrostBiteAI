package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/mattn/go-tty"
)

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	m, _ := ReadMapFromFile("./data/maps/map.txt")

	m.print_height, m.print_width = 20, 65
	m.player.x = 32
	m.player.y = 20

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for true {
		m.print_map()

		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
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
