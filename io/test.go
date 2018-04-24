package main

import (
	"strings"
	"fmt"
)

func main() {
	var s string
	s = "foo,bar,baz,abeer"

	fmt.Println(strings.ContainsAny(s," "))
	fmt.Println(strings.Count(s,"a"))
	fmt.Println(strings.Count(s,"abc"))
	fmt.Println(strings.Fields(s))
	fmt.Println(strings.FieldsFunc("abcdef", func(r rune) bool {
		if r > 'e'{
			return true
		}
		return false

	}))

	fmt.Println(strings.Split(s,","))
	fmt.Println(strings.SplitAfter(s,","))
	fmt.Println(strings.SplitN(s,",",3))

	fmt.Println(strings.HasPrefix(s,"foo"))
	fmt.Println(strings.HasSuffix(s,"er"))

	fmt.Println(strings.Index(s,"er"))
	fmt.Println(strings.IndexFunc(s, func(r rune) bool {

		if r <'u'{
			fmt.Println(string(r))
			return true
		}
		return false
	}))

	fmt.Println(strings.Replace(s,"foo","cpooo",-1))
	r := strings.NewReplacer("foo","baz","a","wwww")
	fmt.Println(r.Replace(s))



}