package main

import "fmt"

func convert(s string, numRows int) string {
    if numRows == 1 {
		return s
	}
	n := 2 * (numRows - 1)
	str := ""
	for j := 0; j < numRows; j++ {
		sj := ""
		for i, c := range s {
			x := (i + numRows - 2) / n
			y := i - x*n
			if y == j || y == -j {
				sj += string(c)
			}
		}
		str += sj
	}
	return str
}
func main() {
    fmt.Println(convert("PAYPALISHIRING", 3))
}