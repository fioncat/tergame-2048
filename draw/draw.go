package draw

import (
	"fmt"
	"strconv"
	"strings"
)

func Area(data [][]int) {
	var width int
	for _, row := range data {
		for _, cell := range row {
			cellStr := strconv.Itoa(cell)
			cellWidth := len(cellStr)
			if cellWidth > width {
				width = cellWidth
			}
		}
	}
	fmtStr := "%" + strconv.Itoa(width) + "d"
	splitLen := len(data[0])*width + len(data[0]) + 1
	splitStr := strings.Repeat("-", splitLen)
	fmt.Println(splitStr)
	for _, row := range data {
		for j, cell := range row {
			if j == 0 {
				fmt.Print("|")
			}
			numStr := fmt.Sprintf(fmtStr, cell)
			fmt.Print(colorNum(cell, numStr) + "|")
		}
		fmt.Println()
		fmt.Println(splitStr)
	}
}

func Points(pt, step int) {
	fmt.Printf("Points: %d, Steps: %d\n", pt, step)
}

func Help() {
	fmt.Println("Press H, J, K, L to operate. Press <Esc> to exit.")
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func colorNum(num int, s string) string {
	if num == 0 {
		return s
	}
	if num <= 8 {
		return green(s)
	}
	if num <= 128 {
		return cyan(s)
	}
	if num <= 1024 {
		return yellow(s)
	}
	if num < 2048 {
		return blue(s)
	}
	return red(s)
}

func green(s string) string {
	return fmt.Sprintf("\033[32;1m%s\033[0m", s)
}

func yellow(s string) string {
	return fmt.Sprintf("\033[33;1m%s\033[0m", s)
}

func cyan(s string) string {
	return fmt.Sprintf("\033[36;1m%s\033[0m", s)
}

func blue(s string) string {
	return fmt.Sprintf("\033[34;1m%s\033[0m", s)
}

func red(s string) string {
	return fmt.Sprintf("\033[31;1m%s\033[0m", s)
}
