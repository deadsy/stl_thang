package stl

import (
	"fmt"
	"testing"
)

func Test_STL_Read(t *testing.T) {

	file_set := []string{
		"../test_data/simple_ascii.stl",
		"../test_data/complex_ascii.stl",
	}

	for _, fname := range file_set {
		solid, err := ReadFile(fname)
		if err != nil {
			t.Error("FAIL")
		}
		fmt.Printf("%s: %d triangles\n", fname, len(solid.Triangles))
	}
}
