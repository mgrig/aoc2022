package spaceleft

import (
	"fmt"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	return &file{
		name: name,
		size: size,
	}
}

// ***** dir

type dir struct {
	name    string
	files   []*file
	subdirs map[string]*dir // name > *dir
}

func newEmptyDir(name string) *dir {
	return &dir{
		name:    name,
		files:   make([]*file, 0),
		subdirs: make(map[string]*dir),
	}
}

func (d dir) tree(level int) {
	fmt.Printf("%s%s\n", strings.Repeat("  ", level), d.name)
	for _, subdir := range d.subdirs {
		subdir.tree(level + 1)
	}
	for _, file := range d.files {
		fmt.Printf("%s%s (file, size=%d)\n", strings.Repeat("  ", level+1), file.name, file.size)
	}
}

func (d dir) size() int {
	size := 0
	for _, file := range d.files {
		size += file.size
	}
	for _, subdir := range d.subdirs {
		size += subdir.size()
	}
	return size
}

// *************

func printPath(currentPath []*dir) string {
	ret := ""
	for i, dir := range currentPath {
		if i > 0 { // skip root
			ret += fmt.Sprintf("/%s", dir.name)
		}
	}
	return ret
}

func parseTree(lines []string) *dir {
	rootDir := newEmptyDir("/")

	// parser states: idle, ls
	// parserState := "idle"

	currentPath := make([]*dir, 1)
	currentPath[0] = rootDir

	for _, line := range lines {
		if strings.HasPrefix(line, "$ ") {
			// is command
			line = line[2:]
			if strings.HasPrefix(line, "cd ") {
				// parserState = "idle"
				// is one of changedir commands...
				line = line[3:]
				if line == ".." {
					// one up
					if len(currentPath) == 0 {
						panic("no current dir, cannot go up")
					}
					if len(currentPath) == 1 {
						// I am in root > nop
						fmt.Println("WARN already in root, cannot go up")
					}
					currentPath = currentPath[0 : len(currentPath)-1]
					fmt.Printf("one up, path: %s\n", printPath(currentPath))
				} else if line == "/" {
					currentPath = make([]*dir, 1)
					currentPath[0] = rootDir
				} else {
					// change to subdir
					currentDir := currentPath[len(currentPath)-1]
					subdir, exists := currentDir.subdirs[line]
					if !exists {
						panic(fmt.Sprintf("missing subdir %s", line))
					}
					currentPath = append(currentPath, subdir)
					fmt.Printf("entered %s, path: %s\n", line, printPath(currentPath))
				}
			} else if line == "ls" {
				// list
				// parserState = "ls"
				// nop
			} else {
				panic("unknown command " + line)
			}
		} else {
			// should be during a ls
			if strings.HasPrefix(line, "dir ") {
				// found a dir > add a new empty dir child
				line = line[4:]
				currentDir := currentPath[len(currentPath)-1]
				subdir := newEmptyDir(line)
				currentDir.subdirs[line] = subdir
			} else {
				// found a file > add it
				currentDir := currentPath[len(currentPath)-1]
				tokens := strings.Split(line, " ")
				if len(tokens) != 2 {
					panic("wrong file line " + line)
				}
				size, err := strconv.Atoi(tokens[0])
				if err != nil {
					panic(err)
				}
				currentDir.files = append(currentDir.files, newFile(tokens[1], size))
			}
		}
	}
	return rootDir
}

func ParseAndShowTree(lines []string) {
	rootDir := parseTree(lines)
	rootDir.tree(0)
}

func TotalSizeDirsBelow100k(lines []string) int {
	rootDir := parseTree(lines)
	var total int
	sizeBelow100k(rootDir, &total)
	return total
}

func sizeBelow100k(d *dir, acc *int) {
	size := 0
	for _, file := range d.files {
		size += file.size
	}
	for _, subdir := range d.subdirs {
		size += subdir.size()
		sizeBelow100k(subdir, acc)
	}

	if size <= 100000 {
		*acc += size
	}
}

func DeleteOneDir(lines []string) int {
	rootDir := parseTree(lines)
	rootSize := rootDir.size()
	unused := 70_000_000 - rootSize
	stillNeeded := 30_000_000 - unused
	fmt.Println("still needed:", stillNeeded)

	sizeDirToDelete := rootSize
	findSizeDirToDelete(rootDir, stillNeeded, &sizeDirToDelete)
	return sizeDirToDelete
}

func findSizeDirToDelete(d *dir, stillNeeded int, acc *int) {
	size := 0
	for _, file := range d.files {
		size += file.size
	}
	for _, subdir := range d.subdirs {
		size += subdir.size()
		findSizeDirToDelete(subdir, stillNeeded, acc)
	}

	if size > stillNeeded && (size-stillNeeded) < (*acc-stillNeeded) {
		*acc = size
	}
}
