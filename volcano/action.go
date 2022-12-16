package volcano

import "fmt"

type action interface {
	String() string
	isOpenAction() bool
	isOpen(name string) bool
	isMoveTo(name string) bool
}

var _ action = move{}

// ***

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
	return fmt.Sprintf("move %s -> %s", m.from, m.to)
}

func (m move) isOpenAction() bool {
	return false
}

func (m move) isOpen(name string) bool {
	return false
}

func (m move) isMoveTo(name string) bool {
	return m.to == name
}

// ***

type open struct {
	name string
}

func newOpenAction(name string) action {
	return open{
		name: name,
	}
}

func (o open) String() string {
	return fmt.Sprintf("open %s", o.name)
}

func (o open) isOpenAction() bool {
	return true
}

func (o open) isOpen(name string) bool {
	return o.name == name
}

func (o open) isMoveTo(name string) bool {
	return false
}
