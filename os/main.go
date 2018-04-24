package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
)

func main() {
	fike,err := ioutil.ReadFile("src/go_notes/README.md")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(fike))

	fmt.Println(path.Dir("src/go_notes/README.md"))
	fmt.Println(filepath.Dir("src/go_notes/README.md"))

	fmt.Println(filepath.Base("src/go_notes/README.md"))
	fmt.Println(path.Base("src/go_notes/README.md"))

	fmt.Println(path.IsAbs("src/go_notes/README.md"))
	fmt.Println(filepath.IsAbs("src/go_notes/README.md"))

	fmt.Println(filepath.Abs("src/go_notes/README.md"))
	fmt.Println(path.Join("src/go_notes/README.md","sdfsdf"))


}
