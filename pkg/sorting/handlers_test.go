package sorting

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestChainStringsort1(t *testing.T) {
	config := &SortingConfig{
		reversed: true,
	}

	chain := makeChain(config)

	src := []string{"A", "B", "C", "A"}
	src = chain.process(src)

	assert.Equal(t, []string{"C", "B", "A", "A"}, src)
}

func TestChainStringsort2(t *testing.T) {
	config := &SortingConfig{
		reversed: true,
		unique:   true,
	}

	chain := makeChain(config)

	src := []string{"A", "B", "C", "A"}
	src = chain.process(src)
	assert.Equal(t, []string{"C", "B", "A"}, src)
}

func TestIntsort1(t *testing.T) {
	config := &SortingConfig{
		reversed: true,
		unique:   true,
		numeric:  true,
	}

	chain := makeChain(config)

	src := []string{"2", "31", "-123", "1"}
	src = chain.process(src)
	assert.Equal(t, []string{"31", "2", "1", "-123"}, src)
}
