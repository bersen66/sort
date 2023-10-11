package tests

import (
	"fmt"
	"testing"

	"github.com/bersen66/sort/pkg/fs"
	"github.com/magiconair/properties/assert"
)

const (
	in  = "in.txt"
	out = "out.txt"
)

func TestReadLines(t *testing.T) {
	lines, err := fs.ReadLines(in)

	assert.Equal(t, nil, err)
	for _, str := range lines {
		fmt.Println(str)
	}
}

func TestFlush(t *testing.T) {
	lines, err := fs.ReadLines(in)
	assert.Equal(t, nil, err)

	err = fs.Flush(lines, out)
	assert.Equal(t, nil, err)
}
