package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data := map[string]bool{
		`madam`:             true,
		`cat`:               false,
		`Madam, I'm Adam`:   true,
		`11,411`:            true,
		`11,c4ac11`:         false,
		`Never odd or even`: true,
	}

	for k, v := range data {
		res := isPalindrome(k)
		if res == v {
			fmt.Println("Cool!")
		} else {
			fmt.Printf("expecting [%v] to be [%v] but [%v]\n", k, v, res)
		}
	}
}

func isPalindrome(s string) bool {
	re := regexp.MustCompile("[^a-zA-Z0-9]")
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	l := len(s)

	for i, n := 0, l/2; i < n; i++ {
		if s[i:i+1] != s[l-i-1:l-i] {
			return false
		}
	}
	return true
}
