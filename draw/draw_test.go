package draw

import "testing"

func TestArea(t *testing.T) {
	data := [][]int{
		{1, 2, 3, 4},
		{1, 200, 3, 4},
		{1, 2, 3, 4},
		{10, 2, 3, 4},
	}

	Area(data)
}
