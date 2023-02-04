package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.Join(listFiles("testdata"), " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := ioutil.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
