package sorting

import (
	"cmp"
	"fmt"
	"os"
	"strconv"
)

func strless(lhs, rhs string) int {
	return cmp.Compare(lhs, rhs)
}

func strgreater(lhs, rhs string) int {
	return -1 * strless(lhs, rhs)
}

func intless(lhs, rhs int) int {
	return cmp.Compare(lhs, rhs)
}

func intgreater(lhs, rhs int) int {
	return -1 * intless(lhs, rhs)
}

func numericcmp(lhs, rhs string) int {
	a, err := strconv.Atoi(lhs)
	if err != nil {
		fmt.Printf("Expected number, has: %v", lhs)
		os.Exit(1)
	}
	b, err := strconv.Atoi(rhs)
	if err != nil {
		fmt.Printf("Expected number, has: %v", rhs)
		os.Exit(1)
	}
	return cmp.Compare(a, b)
}

func columnCompare(lhs, rhs string) int {
	a := getColumnValue(lhs, c.col, c.sep)
	b := getColumnValue(rhs, c.col, c.sep)

	if c.numeric {
		return numericcmp(a, b)
	}
	return cmp.Compare(a, b)
}
