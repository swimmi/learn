package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	str := "1"
	for i := 1; i < n; i++ {
		str = cas(str)
	}
	return str
}
func cas(s string) string {
	outs := ""
	temp := []rune(s)[0]
	count := 0

	for i, c := range s {
		if c != temp {
			outs += strconv.Itoa(count) + string(temp)
			temp = c
			count = 0
		}
		count++
		if i == len(s)-1 {
			outs += strconv.Itoa(count) + string(temp)
		}
	}
	return outs
}
func main() {
	fmt.Println(countAndSay(5))
}