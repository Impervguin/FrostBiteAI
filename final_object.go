package main

type FinalCharacterObject struct {
	CharacterID int
	IsKiller    bool
	Name        string
	Prof        string
}

type FinalObject struct {
	X, Y  int
	Chars []map[int]FinalCharacterObject
}

func (obj FinalObject) get_x() int {
	return obj.X
}

func (obj FinalObject) get_y() int {
	return obj.Y
}

func (obj FinalObject) action() {
	StartTwoSideInteface(obj)
}