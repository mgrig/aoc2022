package problem25

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnafuToDec(t *testing.T) {
	require.Equal(t, 1, newSnafu("1").toDec())
	require.Equal(t, 2, newSnafu("2").toDec())
	require.Equal(t, 3, newSnafu("1=").toDec())
	require.Equal(t, 4, newSnafu("1-").toDec())
	require.Equal(t, 5, newSnafu("10").toDec())
	require.Equal(t, 6, newSnafu("11").toDec())
	require.Equal(t, 7, newSnafu("12").toDec())
	require.Equal(t, 8, newSnafu("2=").toDec())
	require.Equal(t, 9, newSnafu("2-").toDec())
	require.Equal(t, 10, newSnafu("20").toDec())
	require.Equal(t, 15, newSnafu("1=0").toDec())
	require.Equal(t, 20, newSnafu("1-0").toDec())
	require.Equal(t, 2022, newSnafu("1=11-2").toDec())
	require.Equal(t, 12345, newSnafu("1-0---0").toDec())
	require.Equal(t, 314159265, newSnafu("1121-1110-1=0").toDec())
}

func TestDecToSnafu(t *testing.T) {
	require.Equal(t, "1", fromDec(1).value)
	require.Equal(t, "2", fromDec(2).value)
	require.Equal(t, "1=", fromDec(3).value)
	require.Equal(t, "1-", fromDec(4).value)
	require.Equal(t, "10", fromDec(5).value)
	require.Equal(t, "11", fromDec(6).value)
	require.Equal(t, "12", fromDec(7).value)
	require.Equal(t, "2=", fromDec(8).value)
	require.Equal(t, "2-", fromDec(9).value)
	require.Equal(t, "20", fromDec(10).value)
	require.Equal(t, "1=0", fromDec(15).value)
	require.Equal(t, "1-0", fromDec(20).value)
	require.Equal(t, "1=11-2", fromDec(2022).value)
	require.Equal(t, "1-0---0", fromDec(12345).value)
	require.Equal(t, "1121-1110-1=0", fromDec(314159265).value)
}
