package main

import (
	"encoding/json"
	"fmt"
)

type ColorGroup struct {
	Id     int
	Name   string
	Colors []string
}

func main() {
	p := ColorGroup{Id:1,Name:"jialele",Colors:[]string{"sdfsdf","sdf"}}

	//group := ColorGroup{
	//	ID:     1,
	//	Name:   "Reds",
	//	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	//}
	//
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))


}
