package rules

import (
	"errors"
	"slices"
)

type Turn struct {
	target_cells [3]Position
}

func NewTurn(a, b, c Position) (Turn, error) {
	if a == b || a == c || b == c {
		return Turn{}, errors.New("positions must be different")
	}
	return Turn{[3]Position{a, b, c}}, nil
}

func (turn Turn) Contains(p Position) bool {
	return slices.Contains(turn.target_cells[:], p)
}
