package volcano

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
