package echo

import (
	"fmt"
	"strings"
)

func echo1() {
	array := []string{"foo", "bar", "baz", "quax", "phi", "fi", "pho", "fum"}
	var s, sep string
	for i := 0; i < len(array); i++ {
		s += sep + array[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	array := []string{"foo", "bar", "baz", "quax", "phi", "fi", "pho", "fum"}
	s, sep := "", ""
	for _, arg := range array {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	array := []string{"foo", "bar", "baz", "quax", "phi", "fi", "pho", "fum"}
	fmt.Println(strings.Join(array, " "))
}
