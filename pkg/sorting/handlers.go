package sorting

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

type handler interface {
	process(input []string) []string
	setNext(h handler) handler
	hasNext() bool
}

// Sorts strings in specific order
type stringSorter struct {
	next     handler
	reversed bool
}

func newStringSorter(reversed bool) handler {
	return &stringSorter{
		next:     nil,
		reversed: reversed,
	}
}

func (s *stringSorter) process(input []string) []string {
	if s.reversed {
		slices.SortFunc(input, strgreater)
	} else {
		slices.SortFunc(input, strless)
	}

	if s.hasNext() {
		return s.next.process(input)
	}
	return input
}

func (s *stringSorter) hasNext() bool {
	return s.next != nil
}

func (s *stringSorter) setNext(h handler) handler {
	s.next = h
	return s.next
}

//Sorts by numeric value
type numericSorter struct {
	next     handler
	reversed bool
}

func newNumericSorter(reversed bool) handler {
	return &numericSorter{
		reversed: reversed,
		next:     nil,
	}
}

func (n *numericSorter) process(input []string) []string {
	if n.reversed {
		slices.SortFunc(input, func(lhs, rhs string) int {
			return -1 * numericcmp(lhs, rhs)
		})
	} else {
		slices.SortFunc(input, func(lhs, rhs string) int {
			return numericcmp(lhs, rhs)
		})
	}

	if n.hasNext() {
		input = n.next.process(input)
	}

	return input
}

func (n *numericSorter) setNext(next handler) handler {
	n.next = next
	return n.next
}

func (n *numericSorter) hasNext() bool {
	return n != nil
}

// Filters only unique values
type uniqueFilter struct {
	next handler
}

func newUniqueFilter() handler {
	return &uniqueFilter{
		next: nil,
	}
}

func (u *uniqueFilter) process(input []string) []string {
	visited := make(map[string]bool, len(input))
	result := make([]string, 0, cap(input))

	for _, str := range input {
		if _, has := visited[str]; has == false {
			result = append(result, str)
			visited[str] = true
		}
	}

	return result
}

func (u *uniqueFilter) setNext(next handler) handler {
	u.next = next
	return u.next
}

func (u *uniqueFilter) hasNext() bool {
	return u.next != nil
}

// Sorts rows by column value
type columnSorter struct {
	next     handler
	reversed bool
	numeric  bool
	col      int32
	sep      string
}

func newColumnSorter(reversed bool, numeric bool, col int32, sep string) handler {
	return &columnSorter{
		next:     nil,
		reversed: reversed,
		numeric:  numeric,
		col:      col,
		sep:      sep,
	}
}

func getColumnValue(src string, column int32, sep string) string {
	words := strings.Split(src, sep)

	if len(words) < int(column) {
		fmt.Println("Found formating error: not enough columns")
		os.Exit(1)
	}

	return words[column]
}

func (c *columnSorter) process(input []string) []string {

	columnCmp := func(lhs, rhs string) int {
		a := getColumnValue(lhs, c.col, c.sep)
		b := getColumnValue(rhs, c.col, c.sep)

		if c.numeric {
			return numericcmp(a, b)
		}
		return cmp.Compare(a, b)
	}

	if c.reversed {
		slices.SortFunc(input, func(a, b string) int {
			return -1 * columnCmp(a, b)
		})
	} else {
		slices.SortFunc(input, columnCmp)
	}

	if c.hasNext() {
		input = c.next.process(input)
	}

	return input
}

func (c *columnSorter) setNext(next handler) handler {
	c.next = next
	return c.next
}

func (c *columnSorter) hasNext() bool {
	return c.next != nil
}

// Construct chain of handlers
func makeChain(config *SortingConfig) (front handler) {

	if config.byColumn {
		front = newColumnSorter(config.reversed, config.numeric, config.columnNum, config.separator)
	} else if config.numeric {
		front = newNumericSorter(config.reversed)
	} else {
		front = newStringSorter(config.reversed)
	}

	var tail = front

	if config.unique {
		tail = tail.setNext(newUniqueFilter())
	}

	return front
}
