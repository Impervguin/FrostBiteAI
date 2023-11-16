package main

import (
	"fmt"
	"os"
	"os/exec"
	"./interna/obj_map"
)

func main()  {
	m := Map{}
	// m.mat = make([][]char, 10)
	// for i := 0; i < 10; i++ {
	// 	m.mat[i] = make([]char, 10)
	// }
	m.print_height, m.print_width = 7, 7
	m.player.x = 3
	m.player.y = 3
	m.width, m.height = 10, 10
	m.mat = [][]byte{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}

	for true {
		m.print_map()

		var a byte
		fmt.Scanf("%c\n", &a)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		if (a == 'd') {
			m.player.x++
		}
		if (a == 's') {
			m.player.y++
		}
		if (a == 'w') {
			m.player.y--
		}
		if (a == 'a') {
			m.player.x--
		}
			
	}

}
