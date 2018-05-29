package main

import "fmt"

func main() {
	data := map[string]bool{
		"(5+6)*(7+8)/(4+3)":         true,
		"[(5+6)∗(7+8)/(4+3)":        false,
		"(5+6[)∗(7+8])/(4+3)":       false,
		"(45+6)∗(7+8)/(4+3)":        true,
		"(45+6)∗{(7+8)/(4+3)}":      true,
		"(45+6)∗{(7+8)/[(4+3)*47]}": true,
		"[({})({})]":                true,
		"[({)()]":                   false,
	}

	for k, v := range data {
		res := bracesMatch(k)
		if res == v {
			fmt.Println("Cool!")
		} else {
			fmt.Printf("expecting [%v] to be [%v] but [%v]\n", k, v, res)
		}
	}
}

func bracesMatch(s string) bool {
	b := map[string]string{
		"]": "[",
		")": "(",
		"}": "{",
	}
	k := []string{}
	for i, n := 0, len(s); i < n; i++ {
		c := s[i : i+1]
		if c == "{" || c == "[" || c == "(" {
			k = append(k, c)
		} else {
			if c == "}" || c == "]" || c == ")" {
				p := k[len(k)-1 : len(k)]
				k = k[0 : len(k)-1]
				if b[c] != p[0] {
					return false
				}
			}
		}
	}
	return len(k) == 0
}
