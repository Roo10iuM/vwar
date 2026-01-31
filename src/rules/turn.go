package rules

import "slices"

type Turn struct {
	target_cells [3]Position
}

func NewTurn(a, b, c Position) Turn {
	return Turn{[3]Position{a, b, c}}
}

func (turn Turn) Contains(p Position) bool {
	return slices.Contains(turn.target_cells[:], p)
}
