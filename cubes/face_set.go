package cubes

import (
	"fmt"
	"strings"
)

type faceSet struct {
	faces map[face]bool
}

func newFaceSet() faceSet {
	return faceSet{
		faces: make(map[face]bool),
	}
}

func (fs *faceSet) addFace(f face) {
	if f.faceIndex == 0 {
		panic(fmt.Sprintf("wrong face index %d", f.faceIndex))
	}
	fs.faces[f] = true
}

func (fs *faceSet) removeFace(f face) {
	delete(fs.faces, f)
}

func (fs faceSet) contains(f face) bool {
	_, exists := fs.faces[f]
	return exists
}

func (fs faceSet) getAny() face {
	for f, _ := range fs.faces {
		return f
	}
	panic("no face found")
}

func (fs faceSet) String() string {
	strs := make([]string, 0)
	for f := range fs.faces {
		strs = append(strs, f.String())
	}
	return fmt.Sprintf("(%d): %s", len(fs.faces), strings.Join(strs, ", "))
}
