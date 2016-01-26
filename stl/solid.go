package stl

type Solid struct {
	Name      string
	Triangles []Triangle
	// True, if this Solid was read from an ASCII file, and false, if read from a binary file.
	// Used to determine the format when writing to a file.
	IsAscii bool
}
