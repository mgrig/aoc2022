package volcano

import "fmt"

type action interface {
	String() string
	isOpenAction() bool
	isOpen(name string) bool
	isMoveFrom(name string) bool
	isMoveTo(name string) bool
	canApplyTo(name string) bool
	isNop() bool
}

var _ action = move{}
var _ action = open{}
var _ action = nop{}

func equalActions(x, y []action) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// ***

type nop struct{}

func newNop() action {
	return nop{}
}

func (n nop) String() string {
	return "nop"
}

func (n nop) isOpenAction() bool {
	return false
}

func (n nop) isOpen(name string) bool {
	return false
}

func (n nop) isMoveFrom(name string) bool {
	return false
}

func (n nop) isMoveTo(name string) bool {
	return false
}

func (n nop) canApplyTo(name string) bool {
	return true
}

func (n nop) isNop() bool {
	return true
}

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

func (o open) isMoveFrom(name string) bool {
	return false
}

func (o open) isMoveTo(name string) bool {
	return false
}

func (o open) canApplyTo(name string) bool {
	return o.name == name
}

func (o open) isNop() bool {
	return false
}
