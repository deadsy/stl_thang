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
	Name         string
	Triangles    []Triangle
	IsAscii      bool   // true if read from an ASCII file
	BinaryHeader []byte // only used in binary format
}
