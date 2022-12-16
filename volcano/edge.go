package volcano

import "fmt"

type edge struct {
	from, to *node
	weight   int
}

func (e edge) String() string {
	return fmt.Sprintf("(%s - %d - %s)", e.from.name, e.weight, e.to.name)
}
