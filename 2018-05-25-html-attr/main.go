package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data := map[string]string{
		`<tag attr1="v1" attr2="v2"></tag>`:            `v2`,
		`<tag attr1="v1" attr2="v2">some string</tag>`: `v2`,
		`<tag attr1="v1" attr2="v2" />`:                `v2`,
		`<tag attr1="v1" attr2="v2">`:                  ``,
		`<tag attr1="v1" attr2="v2></tag>`:             ``,
		`<tag attr1="v1" attr2="v2></tag" />`:          `v2></tag`,
		`<tag attr1="v1" attr3="v2"></tag>`:            ``,
		`<tag attr1="v1" attr2="v2" attr2="v3"></tag>`: `v2`,
		`<tag attr1="v1" Attr2="V2"></tag>`:            `V2`,
	}

	for k, v := range data {
		res := attr(k, "attr2")
		if res == v {
			fmt.Println("Cool!")
		} else {
			fmt.Printf("expecting [%v] to be [%v] but [%v]\n", k, v, res)
		}
	}
}

func attr(h, a string) string {
	a = strings.ToLower(a)
	r := regexp.MustCompile(`<([^\s]+)(.*)>.*</([^\s>]+)>`)
	m := r.FindStringSubmatch(h)
	if len(m) < 3 {
		r = regexp.MustCompile(`<([^\s]+)(.*)/>`)
		m = r.FindStringSubmatch(h)
		if len(m) < 2 {
			return ""
		}
	} else {
		if m[1] != m[len(m)-1] {
			return ""
		}
	}
	attrList := m[2]
	r = regexp.MustCompile(`\s+([^=]+)="([^"]*)"`)
	m2 := r.FindAllStringSubmatch(attrList, -1)

	for i, n := 0, len(m2); i < n; i++ {
		if strings.ToLower(m2[i][1]) == a {
			return m2[i][2]
		}
	}
	return ""
}
