package volcano

import (
	"fmt"
)

type node struct {
	name string
	flow int
}

func (n node) String() string {
	return fmt.Sprintf("(%s:%d)", n.name, n.flow)
}
