package game

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const (
	directUp = iota
	directDown
	directLeft
	directRight
)

type area struct {
	data    [][]int
	options []int

	nRows int
	nCols int

	point int
}

func (a *area) canMerge() bool {
	for _, row := range a.data {
		if canMerge(row) {
			return true
		}
	}
	for colIdx := 0; colIdx < a.nCols; colIdx++ {
		col := make([]int, a.nRows)
		for rowIdx := 0; rowIdx < a.nRows; rowIdx++ {
			col[rowIdx] = a.data[rowIdx][colIdx]
		}
		if canMerge(col) {
			return true
		}
	}
	return false
}

type pos struct{ i, j int }

func (a *area) fill(n int) {
	var zeroPosList []pos
	for rowIdx := 0; rowIdx < a.nRows; rowIdx++ {
		for colIdx := 0; colIdx < a.nCols; colIdx++ {
			num := a.data[rowIdx][colIdx]
			if num == 0 {
				zeroPosList = append(zeroPosList, pos{
					i: rowIdx, j: colIdx,
				})
			}
		}
	}
	if len(zeroPosList) == 0 {
		return
	}
	for i := 0; i < n; i++ {
		if len(zeroPosList) == 0 {
			return
		}
		posIdx := rand.Intn(len(zeroPosList))
		numIdx := rand.Intn(len(a.options))
		pos := zeroPosList[posIdx]
		num := a.options[numIdx]

		a.data[pos.i][pos.j] = num

		if posIdx == 0 {
			zeroPosList = zeroPosList[1:]
		} else if posIdx == len(zeroPosList)-1 {
			zeroPosList = zeroPosList[:len(zeroPosList)-1]
		} else {
			head := zeroPosList[:posIdx]
			tail := zeroPosList[posIdx+1:]
			zeroPosList = append(head, tail...)
		}
	}
}

func (a *area) mergeAll(direct int) bool {
	switch direct {
	case directRight:
		return a.mergeRows(false)

	case directLeft:
		return a.mergeRows(true)

	case directDown:
		return a.mergeCols(false)

	case directUp:
		return a.mergeCols(true)

	default:
		panic("unknown direct:" + strconv.Itoa(direct))
	}
}

func (a *area) mergeRows(reverse bool) (merged bool) {
	for i, row := range a.data {
		r, ok := a.mergeLine(row, reverse)
		if ok {
			merged = true
		}
		a.data[i] = r
	}
	return
}

func (a *area) mergeCols(reverse bool) (merged bool) {
	for colIdx := 0; colIdx < a.nCols; colIdx++ {
		col := make([]int, a.nRows)
		for rowIdx := 0; rowIdx < a.nRows; rowIdx++ {
			col[rowIdx] = a.data[rowIdx][colIdx]
		}
		r, ok := a.mergeLine(col, reverse)
		if ok {
			merged = true
		}
		for rowIdx := 0; rowIdx < a.nRows; rowIdx++ {
			a.data[rowIdx][colIdx] = r[rowIdx]
		}
	}
	return
}

func canMerge(line []int) bool {
	cursor := -1
	for i, num := range line {
		if num == 0 {
			return true
		}
		if cursor < 0 {
			cursor = i
			continue
		}
		cnum := line[cursor]
		if cnum == num {
			return true
		}
		cursor = i
	}
	return false
}

func (a *area) mergeLine(line []int, reverse bool) ([]int, bool) {
	ori := copySlice(line)
	cursor := -1
	for i, num := range line {
		if num == 0 {
			continue
		}
		if cursor < 0 {
			cursor = i
			continue
		}
		cnum := line[cursor]
		if cnum == num {
			line[i] = num * 2
			line[cursor] = 0
			cursor = -1
			a.point += num
			continue
		}
		cursor = i
	}
	r := make([]int, 0, len(line))
	zeroCnt := 0
	for _, num := range line {
		if num == 0 {
			zeroCnt++
			continue
		}
		r = append(r, num)
	}
	if zeroCnt > 0 {
		zeroList := make([]int, zeroCnt)
		if reverse {
			r = append(r, zeroList...)
		} else {
			r = append(zeroList, r...)
		}
	}
	return r, isMerged(ori, r)
}

func copySlice(line []int) []int {
	slice := make([]int, len(line))
	for i, e := range line {
		slice[i] = e
	}
	return slice
}

func isMerged(origin, result []int) bool {
	for i, e := range origin {
		oe := result[i]
		if e != oe {
			return true
		}
	}
	return false
}
