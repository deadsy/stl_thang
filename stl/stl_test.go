package stl

import (
	"fmt"
	"testing"
)

func Test_STL_Read_Write(t *testing.T) {

	file_set := []string{
		"../test_data/simple_ascii.stl",
		"../test_data/complex_ascii.stl",
		"../test_data/simple_bin.stl",
		"../test_data/complex_bin.stl",
	}

	for _, fname := range file_set {
		solid, err := ReadFile(fname)
		if err != nil {
			t.Error("FAIL")
		}
		fmt.Printf("%s: %d triangles\n", fname, len(solid.Triangles))

		err = solid.WriteFile("stuff.stl")
		if err != nil {
			t.Error("FAIL")
		}
	}
}
