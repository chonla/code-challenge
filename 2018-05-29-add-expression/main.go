package main

import (
	"fmt"
	"strings"
)

func main() {
	data := map[string]string{
		"1+2":            "3",
		"2+34":           "36",
		"1234+5678":      "6912",
		"1234+5678+9101": "16013",
		"1234567890+10":  "1234567900",
		"99999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999+1": "100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	}

	for k, v := range data {
		res := evaluate(k)
		if res == v {
			fmt.Println("Cool!")
		} else {
			fmt.Printf("expecting [%v] to be [%v] but [%v]\n", k, v, res)
		}
	}
}

func evaluate(s string) string {
	exp := strings.Split(s, "+")
	for len(exp) > 1 {
		o1 := exp[0]
		o2 := exp[1]
		exp = exp[2:]

		result := add(o1, o2)
		exp = append(exp, result)
	}

	return exp[0]
}

func add(o1, o2 string) string {
	o1 = reverse(o1)
	o2 = reverse(o2)

	o1l := len(o1)
	o2l := len(o2)

	c := "0"

	result := []string{}

	i := 0
	for i = 0; i < o1l && i < o2l; i++ {
		r, carry := addDigit(o1[i:i+1], o2[i:i+1], c)
		c = carry
		result = append(result, r)
	}

	d := o1[i:]
	if i >= o1l {
		d = o2[i:]
	}

	if len(d) > 0 {
		for i = 0; i < len(d); i++ {
			r, carry := addDigit(d[i:i+1], c, "0")
			c = carry
			result = append(result, r)
		}
	}

	if c != "0" {
		result = append(result, c)
	}

	return reverse(strings.Join(result, ""))
}

func addDigit(d1, d2, c string) (string, string) {
	if d1 == "" {
		d1 = "0"
	}
	if d2 == "" {
		d2 = "0"
	}
	if c == "" {
		c = "0"
	}
	v := (d1[0] - "0"[0]) + (d2[0] - "0"[0]) + (c[0] - "0"[0])
	s := fmt.Sprintf("%02d", v)
	return s[1:2], s[0:1]
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
