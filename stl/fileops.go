package stl

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

// Error when reading binary STL files with incomplete header.
var ErrIncompleteBinaryHeader = errors.New("Incomplete STL binary header, 84 bytes expected")

// Used by ReadFile and ReadAll to signify an incomplete file.
var ErrUnexpectedEOF = errors.New("Unexpected end of file")

var asciiStart = []byte("solid ")

// Extracts an ASCII string from a byte slice. Reads all characters
// from the beginning until a \0 or a non-ASCII character is found.
func extractAsciiString(byteData []byte) string {
	i := 0
	for i < len(byteData) && byteData[i] < byte(128) && byteData[i] != byte(0) {
		i++
	}
	return string(byteData[0:i])
}

// Detects if the file is in STL ASCII format or if it is binary otherwise.
// It will consume 6 bytes and return them.
func isAsciiFile(r io.Reader) (isAscii bool, first6 []byte, err error) {
	first6 = make([]byte, 6) // "solid "
	isAscii = false
	n, readErr := r.Read(first6)
	if n != 6 || readErr == io.EOF {
		err = ErrUnexpectedEOF
		return
	} else if readErr != nil {
		err = readErr
		return
	}

	if bytes.Equal(first6, asciiStart) {
		isAscii = true
	}

	return
}

func ReadAll(r io.Reader) (solid *Solid, err error) {
	isAscii, first6, isAsciiErr := isAsciiFile(r)
	if isAsciiErr != nil {
		err = isAsciiErr
		return
	}

	if isAscii {
		solid, err = readAllAscii(r)
		if solid != nil {
			solid.IsAscii = true
		}
	} else {
		solid, err = readAllBinary(r, first6)
		solid.IsAscii = false
	}

	return
}

func ReadFile(filename string) (solid *Solid, err error) {
	file, openErr := os.Open(filename)
	if openErr != nil {
		err = openErr
		return
	}
	defer file.Close()
	return ReadAll(bufio.NewReader(file))
}
