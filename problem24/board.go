package problem24

type board struct {
	nrRows, nrCols int
	blizzards      [][]string // N, S, E, W, ""
}

func newBoard(nrRows, nrCols int) *board {
	b := board{
		nrRows:    nrRows,
		nrCols:    nrCols,
		blizzards: make([][]string, nrRows),
	}

	for r := 0; r < nrRows; r++ {
		b.blizzards[r] = make([]string, nrCols)
	}

	return &b
}

func (b *board) isOpen(pos, start, end coord, afterStep int) bool {
	if pos == start || pos == end {
		return true
	}
	hitW := b.blizzards[pos.r][(pos.c+afterStep)%b.nrCols] == "W"
	hitE := b.blizzards[pos.r][(pos.c-afterStep%b.nrCols+b.nrCols)%b.nrCols] == "E"
	hitN := b.blizzards[(pos.r+afterStep)%b.nrRows][pos.c] == "N"
	hitS := b.blizzards[(pos.r-afterStep%b.nrRows+b.nrRows)%b.nrRows][pos.c] == "S"

	return !hitW && !hitE && !hitN && !hitS
}
