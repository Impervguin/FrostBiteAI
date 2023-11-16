package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	width, height             int
	print_width, print_height int
	player                    Player
	objs                      []mapObject
	mat                       [][]mapItem
}

func (m Map) print_map() int {
	center_x, center_y := m.player.x, m.player.y
	lx, ly := center_x-m.print_width/2, center_y-m.print_height/2
	rx, ry := center_x+m.print_width/2+1, center_y+m.print_height/2+1
	if m.print_width >= m.width {
		lx = 0
		rx = m.width
	} else {
		for lx < 0 {
			lx++
			rx++
		}
		for rx > m.width {
			rx--
			lx--
		}
	}

	if m.print_height >= m.height {
		ly = 0
		ry = m.height
	} else {
		for ly < 0 {
			ly++
			ry++
		}
		for ry > m.height {
			ry--
			ly--
		}
	}

	for i := ly; i < ry; i++ {
		for j := lx; j < rx; j++ {
			if i == center_y && j == center_x {
				fmt.Printf("%c", 'P')
			} else {
				PrintMapItem(m.mat[i][j])
			}
		}
		fmt.Println()
	}
	return 0
}

func readMapItemMtrFromFile(filename string) ([][]mapItem, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		line := scanner.Text()
		dimensions := strings.Fields(line)
		if len(dimensions) != 2 {
			return nil, 0, 0, fmt.Errorf("некорректный формат размеров матрицы")
		}

		n, err := strconv.Atoi(dimensions[0])
		if err != nil {
			return nil, 0, 0, fmt.Errorf("некорректный формат размера матрицы n")
		}

		m, err := strconv.Atoi(dimensions[1])
		if err != nil {
			return nil, 0, 0, fmt.Errorf("некорректный формат размера матрицы m")
		}

		matrix := make([][]mapItem, n)
		for i, _ := range matrix {
			matrix[i] = make([]mapItem, m)
		}

		for i := 0; i < n && scanner.Scan(); i++ {
			line := scanner.Text()
			for j := 0; j < m && j < len(line); j++ {
				matrix[i][j] = GetMapItem(line[j])
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, 0, 0, err
		}

		return matrix, n, m, nil
	}

	return nil, 0, 0, fmt.Errorf("файл пустой")
}

func ReadMapFromFile(filename string) (Map, error) {
	m := Map{}
	var rc error
	m.mat, m.width, m.height, rc = readMapItemMtrFromFile(filename)
	return m, rc
}
