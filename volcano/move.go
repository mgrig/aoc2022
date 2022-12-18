package volcano

import "fmt"

type move struct {
	from, to string
}

func newMoveAction(from, to string) action {
	if from == to {
		panic("cannot move to the same valve")
	}
	return move{
		from: from,
		to:   to,
	}
}

func (m move) String() string {
	return fmt.Sprintf("(move %s -> %s)", m.from, m.to)
}

func (m move) isOpenAction() bool {
	return false
}

func (m move) isOpen(name string) bool {
	return false
}

func (m move) isMoveFrom(name string) bool {
	return m.from == name
}

func (m move) isMoveTo(name string) bool {
	return m.to == name
}

func (m move) canApplyTo(name string) bool {
	return m.from == name
}

func (m move) isNop() bool {
	return false
}
