package volcano

type nop struct{}

func newNop() action {
	return nop{}
}

func (n nop) String() string {
	return "(nop)"
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
