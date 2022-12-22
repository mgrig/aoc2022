package problem22

import "math"

type coordSystem struct {
	origin   coord
	rotation int // in deg
}

func newCoordSystem(origin coord, rotation int) coordSystem {
	return coordSystem{
		origin:   origin,
		rotation: rotation,
	}
}

func (cs coordSystem) fromOriginToThis(pos coord) coord {
	rotRad := float64(cs.rotation) * math.Pi / 180.0
	cosrot := int(math.Cos(rotRad))
	sinrot := int(math.Sin(rotRad))

	x := pos.r - cs.origin.r
	y := pos.c - cs.origin.c

	newX := x*cosrot + y*sinrot
	newY := y*cosrot - x*sinrot

	return newCoord(newX, newY)
}

func (cs coordSystem) fromThisToOrigin(pos coord) coord {
	rotRad := float64(cs.rotation) * math.Pi / 180.0
	cosrot := int(math.Cos(rotRad))
	sinrot := int(math.Sin(rotRad))

	x := pos.r*cosrot - pos.c*sinrot
	y := pos.c*cosrot + pos.r*sinrot

	newX := x + cs.origin.r
	newY := y + cs.origin.c

	return newCoord(newX, newY)
}
