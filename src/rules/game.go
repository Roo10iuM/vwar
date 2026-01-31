package rules

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
)

type GameState struct {
	cur_turn int
	turns    []Turn
	cells    [][]CellState
	winner   Winner
	width    int
	height   int
}

func NewGame() GameState {
	return NewCustomGame(10, 10)
}

func NewCustomGame(width, height int) GameState {
	cells := make([][]CellState, width)
	for i := range width {
		cells[i] = make([]CellState, height)
	}
	return GameState{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (game_state GameState) Cells() [][]CellState {
	cells := make([][]CellState, game_state.width)
	for i, s := range game_state.cells {
		cells[i] = append(cells[i], s...)
	}
	return cells
}

func (game_state GameState) Winner() Winner {
	return game_state.winner
}

func (game_state GameState) Width() int {
	return game_state.width
}

func (game_state GameState) Height() int {
	return game_state.height
}

func (game_state GameState) IsGameOver() bool {
	return game_state.winner != Unknown
}

func (game_state GameState) Move(turn Turn) (GameState, error) {
	if game_state.cur_turn < 2 {
		if err := game_state.validateFirstTurn(turn); err != nil {
			return game_state, err
		}
		return game_state.move(turn), nil
	}
	if err := game_state.validateTurn(turn); err != nil {
		return game_state, err
	}
	return game_state.move(turn), nil
}

func (game_state GameState) validateTurn(turn Turn) error {
	if game_state.IsGameOver() {
		return errors.New("game over")
	}
	for _, p := range turn.target_cells {
		if !game_state.isOnField(p) {
			return fmt.Errorf(
				"cell (%d; %d) outside field (%d; %d)",
				p.x,
				p.y,
				game_state.width-1,
				game_state.height-1,
			)
		}
		if game_state.isCellOccupied(p) {
			return errors.New("cell occupied")
		}
	}
	for _, p := range turn.target_cells {
		if !game_state.isAlive(p, turn) {
			return fmt.Errorf("cell (%d; %d) is unreacheable", p.x, p.y)
		}
	}
	return nil
}

func (game_state GameState) validateFirstTurn(turn Turn) error {
	if game_state.cur_turn == 0 && !turn.Contains(Position{0, 0}) {
		return errors.New("first move must be to cell (0; 0)")
	}
	if game_state.cur_turn == 1 &&
		!turn.Contains(Position{game_state.width - 1, game_state.height - 1}) {
		return fmt.Errorf(
			"first move must be to cell (%d; %d)",
			game_state.width-1,
			game_state.height-1,
		)
	}
	for _, p := range turn.target_cells {
		if !game_state.isOnField(p) {
			return fmt.Errorf(
				"cell (%d; %d) outside field (%d; %d)",
				p.x,
				p.y,
				game_state.width-1,
				game_state.height-1,
			)
		}
	}
	if game_state.cur_turn == 0 {
		game_state.cells[0][0] = PlayerOneVirus
	}
	if game_state.cur_turn == 1 {
		game_state.cells[game_state.width-1][game_state.height-1] = PlayerTwoVirus
	}
	for _, p := range turn.target_cells {
		if !game_state.isAlive(p, turn) {
			if game_state.cur_turn == 0 {
				game_state.cells[0][0] = Empty
			}
			if game_state.cur_turn == 1 {
				game_state.cells[game_state.width-1][game_state.height-1] = Empty
			}
			return fmt.Errorf("cell (%d; %d) is unreacheable", p.x, p.y)
		}
	}
	if game_state.cur_turn == 0 {
		game_state.cells[0][0] = Empty
	}
	if game_state.cur_turn == 1 {
		game_state.cells[game_state.width-1][game_state.height-1] = Empty
	}
	return nil
}

func (game_state GameState) isAlive(p Position, turn Turn) bool {
	neighbors := game_state.getNeighbors(p)
	new_neigbors := append([]Position{}, neighbors...)
	neighbors_extended := make([]Position, 0)
	for {
		for _, n := range new_neigbors {
			if game_state.cells[n.x][n.y] == game_state.getCurrentPlayerVirus() {
				return true
			} else if game_state.cells[n.x][n.y] == game_state.getCurrentPlayerCell() ||
				turn.Contains(n) {
				neighbors_extended = lo.Union(neighbors_extended, game_state.getNeighbors(n))
			}
		}
		_, new_neigbors = lo.Difference(neighbors, neighbors_extended)
		if len(new_neigbors) == 0 {
			break
		}
		neighbors = append(neighbors, new_neigbors...)
	}
	return false
}

func (game_state GameState) getNeighbors(p Position) []Position {
	neighbors := make([]Position, 0, 5)
	if p.y != game_state.height-1 {
		//neighbor 1
		if p.x != 0 {
			neighbors = append(neighbors, NewPosition(p.x-1, p.y+1))
		}
		//neighbor 2
		neighbors = append(neighbors, NewPosition(p.x, p.y+1))
		//neighbor 3
		if p.x != game_state.width-1 {
			neighbors = append(neighbors, NewPosition(p.x+1, p.y+1))
		}
	}

	//neighbor 4
	if p.x != 0 {
		neighbors = append(neighbors, NewPosition(p.x-1, p.y))
	}
	//neighbor 6
	if p.x != game_state.width-1 {
		neighbors = append(neighbors, NewPosition(p.x+1, p.y))
	}

	if p.y != 0 {
		//neighbor 7
		if p.x != 0 {
			neighbors = append(neighbors, NewPosition(p.x-1, p.y-1))
		}
		//neighbor 8
		neighbors = append(neighbors, NewPosition(p.x, p.y-1))
		//neighbor 9
		if p.x != game_state.width-1 {
			neighbors = append(neighbors, NewPosition(p.x+1, p.y-1))
		}
	}
	return neighbors
}

func (game_state GameState) move(turn Turn) GameState {
	for i := range game_state.width {
		for j := range game_state.height {
			if turn.Contains(Position{i, j}) {
				game_state.cells[i][j] = game_state.toggleCell(Position{i, j})
			}
		}
	}
	game_state.turns = append(game_state.turns, turn)
	game_state.cur_turn += 1
	if game_state.isGameOver() {
		if game_state.cur_turn%2 == 0 {
			game_state.winner = PlayerTwo
		} else {
			game_state.winner = PlayerOne
		}
	}
	return game_state
}

func (game_state GameState) toggleCell(p Position) CellState {
	switch game_state.cells[p.x][p.y] {
	case Empty:
		return game_state.getCurrentPlayerVirus()
	case game_state.getEnemyPlayerVirus():
		return game_state.getEnemyPlayerCell()
	default:
		return Empty
	}
}

func (game_state GameState) getCurrentPlayerVirus() CellState {
	if game_state.cur_turn%2 == 0 {
		return PlayerOneVirus
	} else {
		return PlayerTwoVirus
	}
}

func (game_state GameState) getCurrentPlayerCell() CellState {
	if game_state.cur_turn%2 == 0 {
		return PlayerOneCell
	} else {
		return PlayerTwoCell
	}
}

func (game_state GameState) getEnemyPlayerVirus() CellState {
	if game_state.cur_turn%2 == 0 {
		return PlayerTwoVirus
	} else {
		return PlayerOneVirus
	}
}

func (game_state GameState) getEnemyPlayerCell() CellState {
	if game_state.cur_turn%2 == 0 {
		return PlayerTwoCell
	} else {
		return PlayerOneCell
	}
}

func (game_state GameState) isOnField(p Position) bool {
	return p.x >= 0 &&
		p.x <= game_state.width-1 &&
		p.y >= 0 &&
		p.y <= game_state.height-1
}

func (game_state GameState) isCellOccupied(p Position) bool {
	switch game_state.cells[p.x][p.y] {
	case Empty:
		return false
	case game_state.getEnemyPlayerVirus():
		return false
	default:
		return true
	}
}

func (game_state GameState) isGameOver() bool {
	if game_state.cur_turn < 2 {
		return false
	}
	num_of_possible_movies := 0
	for c := range game_state.height {
		for r := range game_state.width {
			if game_state.isPotential(Position{r, c}) {
				num_of_possible_movies += 1
				if num_of_possible_movies >= 3 {
					return false
				}
			}
		}
	}
	return true
}

func (game_state GameState) isPotential(p Position) bool {
	if game_state.isCellOccupied(p) {
		return false
	}
	neighbors := game_state.getNeighbors(p)
	new_neigbors := append([]Position{}, neighbors...)
	neighbors_extended := make([]Position, 0)
	for {
		for _, n := range new_neigbors {
			if game_state.cells[n.x][n.y] == game_state.getCurrentPlayerVirus() {
				return true
			} else if game_state.cells[n.x][n.y] == Empty ||
				game_state.cells[n.x][n.y] == game_state.getCurrentPlayerCell() ||
				game_state.cells[n.x][n.y] == game_state.getEnemyPlayerVirus() {
				neighbors_extended = lo.Union(neighbors_extended, game_state.getNeighbors(n))
			}
		}
		_, new_neigbors = lo.Difference(neighbors, neighbors_extended)
		if len(new_neigbors) == 0 {
			break
		}
		neighbors = append(neighbors, new_neigbors...)
	}
	return false
}
