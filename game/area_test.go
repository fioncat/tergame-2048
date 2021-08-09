package game

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	lines := [][]int{
		{2, 2, 0, 0},
		{4, 0, 4, 0},
		{2, 0, 0, 2},
		{2, 0, 4, 0},
		{4, 8, 8, 4},
		{8, 8, 8, 8},
		{4, 4, 2, 4},
		{0, 4, 4, 0},
		{0, 4, 8, 8},
		{0, 2, 2, 4},
		{0, 0, 4, 4},
		{4, 4, 0, 0},
	}

	for _, line := range lines {
		merged := merge(line, false)
		fmt.Printf("origin = %v, merged = %v\n",
			line, merged)
	}
}
