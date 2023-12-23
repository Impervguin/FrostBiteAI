package main

import (
	"errors"
)

type Player struct {
	x, y int
}

func (m *Map) PlayerMoveUp() (int, error) {
	return m.PlayerMove(m.player.x, m.player.y-1)
}

func (m *Map) PlayerMoveRight() (int, error) {
	return m.PlayerMove(m.player.x+1, m.player.y)
}

func (m *Map) PlayerMoveLeft() (int, error) {
	return m.PlayerMove(m.player.x-1, m.player.y)
}

func (m *Map) PlayerMoveDown() (int, error) {
	return m.PlayerMove(m.player.x, m.player.y+1)
}

func (m *Map) PlayerMove(x, y int) (int, error) {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return 0, errors.New("Movement out of bounds.")
	}
	for _, obj := range m.objs {
		if obj.get_x() == x && obj.get_y() == y {
			return obj.action(), nil
		}
	}

	if m.mat[y][x].passable {
		m.player.x = x
		m.player.y = y
	}
	return 0, errors.New("Not passable object.")
}
