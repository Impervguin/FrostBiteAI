package map_obj

import (
	"fmt"
)

type Map struct {
	width, height             int
	print_width, print_height int
	player                    Player
	objs                      []mapObject
	mat                       [][]byte
}

func (m Map) print_map() int {
	center_x, center_y := m.player.x, m.player.y;
	lx, ly := center_x - m.print_width / 2, center_y - m.print_height / 2
	rx, ry := center_x + m.print_width / 2 + 1, center_y + m.print_height / 2 + 1
	if (m.print_width >= m.width) {
		lx = 0
		rx = m.width
	} else {
		for (lx < 0) {
			lx++;
			rx++;
		}
		for (rx > m.width) {
			rx--;
			lx--;
		}
	}

	if (m.print_height >= m.height) {
		ly = 0
		ry = m.height
	} else {
		for (ly < 0) {
			ly++;
			ry++;
		}
		for (ry > m.height) {
			ry--;
			ly--;
		}
	}

	for i := ly; i < ry; i++ {
		for j := lx; j < rx; j++ {
			if (i == center_y && j == center_x) {
				fmt.Printf("%c", 'P')
			} else {
				fmt.Printf("%c", m.mat[i][j])
			}
		}
		fmt.Println()
	}
	return 0
}
