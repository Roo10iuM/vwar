package vwar

import (
	"fmt"
	"testing"

	"github.com/roo10ium/vwar/rules"
)

func printField(game rules.GameState) {
	for c := range game.Height() {
		for r := range game.Width() {
			fmt.Print(game.Cells()[r][c])
		}
		fmt.Println()
	}
}

func Test4x4(t *testing.T) {
	game := rules.NewCustomGame(4, 4)
	turn_1, _ := rules.NewTurn(
		rules.NewPosition(0, 0),
		rules.NewPosition(1, 0),
		rules.NewPosition(2, 0),
	)
	game, err := game.Move(turn_1)

	turn_2, _ := rules.NewTurn(
		rules.NewPosition(3, 3),
		rules.NewPosition(2, 2),
		rules.NewPosition(1, 1),
	)
	game, err = game.Move(turn_2)

	turn_3, _ := rules.NewTurn(
		rules.NewPosition(3, 3),
		rules.NewPosition(2, 2),
		rules.NewPosition(1, 1),
	)
	game, err = game.Move(turn_3)
	fmt.Println(err)
	printField(game)
	fmt.Println(game.IsGameOver(), game.Winner())
}
