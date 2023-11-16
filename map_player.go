package main

import (
	"errors"
)

type Player struct {
	x, y int
}

func (m *Map) PlayerMoveUp() error {
	return m.PlayerMove(m.player.x, m.player.y-1)
}

func (m *Map) PlayerMoveRight() error {
	return m.PlayerMove(m.player.x+1, m.player.y)
}

func (m *Map) PlayerMoveLeft() error {
	return m.PlayerMove(m.player.x-1, m.player.y)
}

func (m *Map) PlayerMoveDown() error {
	return m.PlayerMove(m.player.x, m.player.y+1)
}

func (m *Map) PlayerMove(x, y int) error {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return errors.New("Movement out of bounds.")
	}
	for _, obj := range m.objs {
		if obj.get_x() == x && obj.get_y() == y {
			obj.action()
			return nil
		}
	}

	// fmt.Println(m.mat[y][x])
	if m.mat[y][x].passable {
		m.player.x = x
		m.player.y = y
	}
	return errors.New("Not passable object.")
}
