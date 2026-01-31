package rules

import (
	"errors"
)

type Winner int
type CellState int

const (
	Empty CellState = iota
	PlayerOneVirus
	PlayerTwoVirus
	PlayerOneVirusKilled
	PlayerTwoVirusKilled
)

const (
	PlayerOne Winner = iota
	PlayerTwo
	Unknown
)

type GameState struct {
	cur_turn      int
	turns         []Turn
	cells         [][]CellState
	winner        Winner
	width, height int
}

func (game_state GameState) Move(turn Turn) (GameState, error) {
	if game_state.IsGameOver() {
		return game_state, errors.New("game over")
	}
	if !game_state.isTurnValide(turn) {
		return game_state, errors.New("invalid move")
	}
	return game_state.makeMove(turn), nil
}

func (game_state GameState) IsGameOver() bool {
	return game_state.winner != Unknown
}

func (game_state GameState) isTurnValide(turn Turn) bool {
	if game_state.cur_turn == 0 && !turn.Contains(Position{0, 0}) {
		return false
	}
	if game_state.cur_turn == 1 && !turn.Contains(Position{game_state.width, game_state.height}) {
		return false
	}
	for _, p := range turn.target_cells {
		if !game_state.isAlive(p) {
			return false
		}
	}
	return true
}

func (game_state GameState) isAlive(p Position) bool {
	return true
}

func (game_state GameState) makeMove(turn Turn) GameState {
	return game_state
}

// func (game_state GameState) Move(turn Turn) error {
// 	game_state.cur_turn = turn
// 	if !game_state.isCurTurnValide() {
// 		return errors.New("invalid move")
// 	}
// 	game_state.setCells(turn)
// 	// TODO
// 	return nil
// }

// func (game_state GameState) isCurTurnValide() bool {
// 	for _, target_cell := range game_state.cur_turn.target_cells {
// 		if !game_state.isAlive(target_cell) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (game_state GameState) isAlive(cell Cell) bool {
// 	var chain []Cell
// 	for _, neighbour := range cell.getNeighbours() {
// 		chain = append(chain, neighbour)
// 	}
// 	//TODO
// }

// func (game_state GameState) getNeighbours(cell Cell) []Cell {
// 	var neighbours []Cell
// 	row, column := cell.position.row, cell.position.column
// 	if row != 9 {
// 		if column != 0 {
// 			neighbours = append(neighbours, game_state.cells[row+1][column-1])
// 		}
// 		neighbours = append(neighbours)
// 	}
// 	// TODO
// 	return neighbours
// }

// func (game_state GameState) setCells(turn Turn) error {
// 	//TODO
// 	return nil
// }

// func (game_state GameState) toggleCell(target_postion Position) error {
// 	if game_state.isKilledCell(target_postion) {
// 		return fmt.Errorf(
// 			"do not disturb the dead (%v; %v)",
// 			target_postion.row,
// 			target_postion.column,
// 		)
// 	}
// 	if game_state.isFriendlyCell(target_postion) {
// 		return fmt.Errorf(
// 			"friendly fire is prohibited (%v; %v)",
// 			target_postion.row,
// 			target_postion.column,
// 		)
// 	}
// 	cell := game_state.cells[target_postion.row][target_postion.column]
// 	switch cell.state {
// 	case Empty:
// 		cell.state = game_state.getCurrentPlayerVirus()
// 	case game_state.getEnemyPlayerVirus():
// 		cell.state = game_state.getEnemyPlayerVirusKilled()
// 	}
// 	return nil
// }

// func (game_state GameState) isKilledCell(target_postion Position) bool {
// 	cell := game_state.cells[target_postion.row][target_postion.column]
// 	if cell.state == PlayerOneVirusKilled || cell.state == PlayerTwoVirusKilled {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func (game_state GameState) isFriendlyCell(target_postion Position) bool {
// 	cell := game_state.cells[target_postion.row][target_postion.column]
// 	switch cell.state {
// 	case game_state.getCurrentPlayerVirus():
// 		return true
// 	case game_state.getEnemyPlayerVirusKilled():
// 		return true
// 	default:
// 		return false
// 	}
// }

// func (game_state GameState) getCurrentPlayerVirus() CellState {
// 	if len(game_state.turns)%2 == 0 {
// 		return PlayerOneVirus
// 	} else {
// 		return PlayerTwoVirus
// 	}
// }

// func (game_state GameState) getEnemyPlayerVirus() CellState {
// 	if len(game_state.turns)%2 == 0 {
// 		return PlayerTwoVirus
// 	} else {
// 		return PlayerOneVirus
// 	}
// }

// func (game_state GameState) getEnemyPlayerVirusKilled() CellState {
// 	if len(game_state.turns)%2 == 0 {
// 		return PlayerTwoVirusKilled
// 	} else {
// 		return PlayerOneVirusKilled
// 	}
// }
