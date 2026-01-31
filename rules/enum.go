package rules

type Winner int
type CellState int

const (
	Empty CellState = iota
	PlayerOneVirus
	PlayerTwoVirus
	PlayerOneCell
	PlayerTwoCell
)

const (
	Unknown Winner = iota
	PlayerOne
	PlayerTwo
)

// func (cs CellState) String() string {
// 	return [...]string{
// 		"Empty",
// 		"PlayerOneVirus",
// 		"PlayerTwoVirus",
// 		"PlayerOneCell",
// 		"PlayerTwoCell",
// 	}[cs]
// }

func (w Winner) String() string {
	return [...]string{
		"Unknown",
		"PlayerOne",
		"PlayerTwo",
	}[w]
}
