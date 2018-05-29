package main

import (
	"fmt"
	"strings"
)

func main() {
	data := map[int]string{
		1: `*`,
		2: ` *
***
 *`,
		3: `  *
 ***
*****
 ***
  *`,
		4: `   *
  ***
 *****
*******
 *****
  ***
   *`,
		5: `    *
   ***
  *****
 *******
*********
 *******
  *****
   ***
    *`,
	}

	for k, v := range data {
		res := dia(k)
		if res == v {
			fmt.Println("Cool!")
		} else {
			fmt.Printf("expecting [%v] to be [%v] but [%v]\n", k, v, res)
		}
	}
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
