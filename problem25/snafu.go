package problem25

import (
	"math"
)

type snafu struct {
	value string
}

func newSnafu(snafuValue string) snafu {
	return snafu{
		value: snafuValue,
	}
}

func snafuToInt(run rune) int {
	switch run {
	case '=':
		return -2
	case '-':
		return -1
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	default:
		panic("wrong snafu char")
	}
}

func (s snafu) toDec() int {
	ret := 0
	len := len(s.value)
	for i, run := range s.value {
		powerOf5 := len - i - 1
		ret += snafuToInt(run) * int(math.Pow(5, float64(powerOf5)))
	}
	return ret
}

var SNAFU_CHARS = "012=-"

func fromDec(decValue int) snafu {
	snafuValue := ""
	for decValue > 0 {
		snafuValue = string(SNAFU_CHARS[decValue%5]) + snafuValue

		decValue = decValue - (((decValue + 2) % 5) - 2)
		decValue /= 5
	}
	return newSnafu(snafuValue)
}
