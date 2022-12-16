package volcano

import (
	"fmt"
	"strings"
)

type node struct {
	name    string
	flow    int
	edgesTo []*node
}

func (n node) String() string {
	edgesStr := make([]string, len(n.edgesTo))
	for i, n := range n.edgesTo {
		edgesStr[i] = n.name
	}
	return fmt.Sprintf("%s, %d, %v", n.name, n.flow, "["+strings.Join(edgesStr, ", ")+"]")
}
