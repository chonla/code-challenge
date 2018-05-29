package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(dia(1))
	fmt.Println(dia(2))
	fmt.Println(dia(3))
	fmt.Println(dia(4))
	fmt.Println(dia(5))
	fmt.Println(dia(6))
}

func dia(n int) string {
	lines := []string{}
	for i := 1; i <= n; i++ {
		lines = append(lines, diaLine(n, i, i*2-1))
	}
	for i := n - 1; i > 0; i-- {
		lines = append(lines, diaLine(n, i, i*2-1))
	}
	return strings.Join(lines, "\n")
}

func diaLine(n, s, a int) string {
	o := ""
	for j := n; j > s; j-- {
		o += " "
	}
	for j := 0; j < a; j++ {
		o += "*"
	}
	return o
}
