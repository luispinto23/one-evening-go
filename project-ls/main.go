package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	dirs := listFiles("testdata")

	for _, dir := range dirs {
		fmt.Println(dir)
	}
}

func listFiles(dirname string) []string {
	var dirs []string
	var showAll = flag.Bool("a", false, "show all files (including hidden)")
	flag.Parse()

	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if *showAll {
			dirs = append(dirs, f.Name())
		} else if !strings.HasPrefix(f.Name(), ".") {
			dirs = append(dirs, f.Name())
		}
	}

	return dirs
}
