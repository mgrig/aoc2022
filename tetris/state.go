package tetris

type state struct {
	upcomingShapeIndex int
	upcomingWindIndex  int
	sky                skyline
}

func newState(upcomingShapeIndex, upcomingWindIndex int, sky skyline) state {
	return state{
		upcomingShapeIndex: upcomingShapeIndex,
		upcomingWindIndex:  upcomingWindIndex,
		sky:                sky,
	}
}

// ***

type stateHistory struct {
	history []state
}

func newStateHistory() stateHistory {
	return stateHistory{
		history: make([]state, 0),
	}
}

func (sh *stateHistory) addState(s state) (gotRepetition bool, firstIndex, secondIndex int) {
	for i := range sh.history {
		oldState := sh.history[i]
		if oldState.upcomingShapeIndex != s.upcomingShapeIndex ||
			oldState.upcomingWindIndex != s.upcomingWindIndex ||
			!oldState.sky.normalizedEquals(s.sky) {
			continue
		}
		// fmt.Println("i=", i)
		// fmt.Println("len hist=", len(sh.history))
		// fmt.Println("state", s)
		// s.sky.plot()
		return true, i, len(sh.history)
	}
	sh.history = append(sh.history, s)
	return false, -1, -1
}
