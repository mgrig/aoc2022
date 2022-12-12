package monkey

import (
	"aoc2022/common"
	"fmt"
	"math/big"
)

type WorryLevel interface {
	square() WorryLevel
	multInt(scalar int) WorryLevel
	addInt(scalar int) WorryLevel
	divInt(scalar int) WorryLevel
	divisibleBy(scalar int) bool
}

var _ WorryLevel = BigWorryLevel{}
var _ WorryLevel = IntWorryLevel{}
var _ WorryLevel = U64WorryLevel{}

// ***

var moduloValues = []int{2, 3, 5, 7, 11, 13, 17, 19}

// var moduloValues = []int{13, 17, 19, 23}

type SmartWL struct {
	modulo map[int]int
}

func (s *SmartWL) setValue(value int) {
	for _, v := range moduloValues {
		s.modulo[v] = value % v
	}
}

func newSmartWL(value int) WorryLevel {
	swl := SmartWL{
		modulo: make(map[int]int),
	}
	swl.setValue(value)
	var wl WorryLevel = swl
	return wl
}

func (swl SmartWL) withOperation(fn func(int, int) int) SmartWL {
	ret := SmartWL{}
	ret.modulo = make(map[int]int)
	for k, v := range swl.modulo {
		ret.modulo[k] = fn(k, v)
	}
	return ret
}

func (swl SmartWL) square() WorryLevel {
	return swl.withOperation(func(k, v int) int {
		return (v * v) % k
	})
}

func (swl SmartWL) multInt(scalar int) WorryLevel {
	return swl.withOperation(func(k, v int) int {
		return (v * scalar) % k
	})
}

func (swl SmartWL) addInt(scalar int) WorryLevel {
	return swl.withOperation(func(k, v int) int {
		return (v + scalar) % k
	})
}

func (swl SmartWL) divInt(scalar int) WorryLevel {
	panic("divInt not support for SmartWL")
}

func (swl SmartWL) divisibleBy(scalar int) bool {
	value, exists := swl.modulo[scalar]
	if !exists {
		panic(fmt.Sprintf("missing modulor for %d", scalar))
	}
	return value == 0
}

// ***

type U64WorryLevel struct {
	worryLevel uint64
}

func newU64WorryLevel(value uint64) WorryLevel {
	var wl WorryLevel = U64WorryLevel{
		worryLevel: uint64(value),
	}
	return wl
}

func (u64wl U64WorryLevel) String() string {
	return fmt.Sprintf("%d", u64wl.worryLevel)
}

func (u64wl U64WorryLevel) square() WorryLevel {
	return newU64WorryLevel(u64wl.worryLevel * u64wl.worryLevel)
}

func (u64wl U64WorryLevel) multInt(scalar int) WorryLevel {
	return newU64WorryLevel(u64wl.worryLevel * uint64(scalar))
}

func (u64wl U64WorryLevel) addInt(scalar int) WorryLevel {
	return newU64WorryLevel(u64wl.worryLevel + uint64(scalar))
}

func (u64wl U64WorryLevel) divInt(scalar int) WorryLevel {
	// x := common.Uint64ToBigInt(u64wl.worryLevel)
	// y := big.NewInt(int64(scalar))

	// z := big.NewInt(0)
	// z.Quo(x, y)

	// return newU64WorryLevel(z.Uint64())

	return newU64WorryLevel(u64wl.worryLevel / uint64(scalar))
}

func (u64wl U64WorryLevel) divisibleBy(scalar int) bool {
	// x := common.Uint64ToBigInt(u64wl.worryLevel)
	// y := big.NewInt(int64(scalar))

	// // x := new(big.Int).SetUint64(u64wl.worryLevel)
	// // y := new(big.Int).SetUint64(uint64(scalar))

	// z := big.NewInt(0)
	// z.Rem(x, y)

	// return z.Uint64() == 0

	return u64wl.worryLevel%uint64(scalar) == 0
}

// ***

type BigWorryLevel struct {
	worryLevel big.Int
}

func newBigWorryLevel(value int) WorryLevel {
	var wl WorryLevel = BigWorryLevel{
		worryLevel: *common.IntToBigInt(value),
	}
	return wl
}

func bigIntToWL(value *big.Int) WorryLevel {
	var wl WorryLevel = BigWorryLevel{
		worryLevel: *value,
	}
	return wl
}

func (bwl BigWorryLevel) String() string {
	return fmt.Sprintf("%v", bwl.worryLevel)
}

func (bwl BigWorryLevel) square() WorryLevel {
	z := big.NewInt(0)
	z.Mul(&bwl.worryLevel, &bwl.worryLevel)
	return bigIntToWL(z)
}

func (bwl BigWorryLevel) multInt(scalar int) WorryLevel {
	z := big.NewInt(0)
	z.Mul(&bwl.worryLevel, big.NewInt(int64(scalar)))
	return bigIntToWL(z)
}

func (bwl BigWorryLevel) addInt(scalar int) WorryLevel {
	z := big.NewInt(0)
	z.Add(&bwl.worryLevel, big.NewInt(int64(scalar)))
	return bigIntToWL(z)
}

func (bwl BigWorryLevel) divInt(scalar int) WorryLevel {
	z := big.NewInt(0)
	// z.Div(&bwl.worryLevel, big.NewInt(int64(scalar)))
	z.Quo(&bwl.worryLevel, big.NewInt(int64(scalar)))

	return bigIntToWL(z)
}

func (bwl BigWorryLevel) divisibleBy(scalar int) bool {
	// z := big.NewInt(0)
	z := common.Uint64ToBigInt(uint64(0))
	z.Rem(&bwl.worryLevel, common.Uint64ToBigInt(uint64(scalar)))

	x := bwl.worryLevel.Uint64()
	y := uint64(scalar)
	rem := x % y
	if rem != z.Uint64() {
		fmt.Printf("bwl.worryLevel = %T %v\n", bwl.worryLevel, bwl.worryLevel)
		fmt.Printf("&bwl.worryLevel = %T %v\n", &bwl.worryLevel, &bwl.worryLevel)

		fmt.Printf("big: %v %% %d\n", &bwl.worryLevel, scalar)
		panic(fmt.Sprintf("WTF %d %% %d > big: %d, u64: %d", x, y, z.Uint64(), rem))
	}

	return z.Uint64() == 0
}

// ***

type IntWorryLevel struct {
	worryLevel int
}

func newIntWorryLevel(value int) WorryLevel {
	var wl WorryLevel = IntWorryLevel{
		worryLevel: value,
	}
	return wl
}

func (iwl IntWorryLevel) square() WorryLevel {
	return newIntWorryLevel(iwl.worryLevel * iwl.worryLevel)
}

func (iwl IntWorryLevel) multInt(scalar int) WorryLevel {
	return newIntWorryLevel(iwl.worryLevel * scalar)
}

func (iwl IntWorryLevel) addInt(scalar int) WorryLevel {
	return newIntWorryLevel(iwl.worryLevel + scalar)
}

func (iwl IntWorryLevel) divInt(scalar int) WorryLevel {
	return newIntWorryLevel(iwl.worryLevel / scalar)
}

func (iwl IntWorryLevel) divisibleBy(scalar int) bool {
	return iwl.worryLevel%scalar == 0
}
