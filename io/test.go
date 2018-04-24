package main

import (
	"strings"
	"fmt"
)

func main() {
	var s string
	s = "abcdefj abeer"

	fmt.Println(strings.ContainsAny(s," "))
	fmt.Println(strings.Count(s,"a"))
	fmt.Println(strings.Count(s,"abc"))
	fmt.Println(strings.Fields(s))
	fmt.Println(strings.Count(s,"a"))
	fmt.Println(strings.Count(s,"a"))


}