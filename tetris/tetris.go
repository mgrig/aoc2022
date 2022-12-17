package tetris

import "fmt"

func Part1(line string, totalRocks int) int {
	mapWidth := 7

	bottom := newBottom(mapWidth)
	wind := newWind(line)

	placedRocks := 0
	steps := 0
	history := newStateHistory()
	gotRepetition := false
	deltaHeight := 0
	for { // simulation steps

		if !gotRepetition {
			var firstIndex, secondIndex int
			gotRepetition, firstIndex, secondIndex = history.addState(newState(placedRocks%5, wind.getUpcomingIndex()%len(line), bottom.skyline()))
			if gotRepetition {
				fmt.Printf("1st index: %d\n2nd index: %d\n", firstIndex, secondIndex)
				_, heightAtFirstIndex := history.history[firstIndex].sky.minmax()
				_, heightAtSecondIndex := bottom.skyline().minmax()
				rocksPeriod := secondIndex - firstIndex
				numberRepetitions := (totalRocks - firstIndex) / rocksPeriod
				remainder := (totalRocks - firstIndex) % rocksPeriod
				// bring totalRocks close to the current top,
				// and compute the deltaHeight we have to apply at the end
				totalRocks = secondIndex + remainder
				deltaHeight = (numberRepetitions - 1) * (heightAtSecondIndex - heightAtFirstIndex)
			}
		}

		if placedRocks == totalRocks {
			break
		}

		// spawn rock (type + start position)
		spawnPos := newCoord(2, 4+bottom.highest())
		// fmt.Println("bottom highest", bottom.highest())
		// fmt.Println("spawn pos", spawnPos)
		rock := GetShape(placedRocks, spawnPos)

		for { // falling piece
			// get wind at current step
			w := wind.getNext()
			steps++

			// move piece left & down if possible
			var ok bool
			if w == "<" {
				// fmt.Printf("move left: %s ", rock)
				_, rock = moveLeft(rock, &bottom, mapWidth)
				// fmt.Printf(" %t > %s\n", ok, rock)
			} else if w == ">" {
				// fmt.Printf("move right: %s ", rock)
				_, rock = moveRight(rock, &bottom, mapWidth)
				// fmt.Printf(" %t > %s\n", ok, rock)
			} else {
				panic("bad wind")
			}

			// fmt.Printf("move down: %s ", rock)
			ok, rock = moveDown(rock, &bottom)
			// fmt.Printf(" %t > %s\n", ok, rock)
			if !ok {
				// piece could not move down > at rest
				// fmt.Println(rock, "at rest")

				// add to bottom
				for _, aC := range getAbsoluteShapeCoords(rock) {
					bottom.addCoord(aC)
				}

				break
			}
		}
		placedRocks++
	}

	// bottom.plot(true)

	return bottom.highest() + deltaHeight
}
