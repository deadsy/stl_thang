package stl

import (
	"github.com/deadsy/stl_thang/vec"
)

type Triangle struct {
	Normal     vec.V3
	Vertices   [3]vec.V3
	Attributes uint16
}

type Solid struct {
	Name      string
	Triangles []Triangle
	// True, if this Solid was read from an ASCII file, and false, if read from a binary file.
	// Used to determine the format when writing to a file.
	IsAscii bool
}
