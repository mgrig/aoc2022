package volcano

import "fmt"

type open struct {
	name string
}

func newOpenAction(name string) action {
	return open{
		name: name,
	}
}

func (o open) String() string {
	return fmt.Sprintf("(open %s)", o.name)
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
