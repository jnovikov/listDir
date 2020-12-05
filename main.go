package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	a = flag.Bool("a", false, "Print all info")
)

func hrSize(fsize int64) string {
	return "0KB"
}

func printAll(file os.FileInfo) {
	time := file.ModTime().Format("Jan 06 15:4")
	fSize := strconv.Itoa(int(file.Size()))
	fmt.Printf("%s %s %s \n", fSize, time, file.Name())
}

func main() {
	flag.Parse()
	dir := "."
	if flag.NArg() >= 1 {
		dir = flag.Arg(0)
	}
	files, _ := ioutil.ReadDir(dir)

	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			fmt.Println(file.Name())
		}
	}
}
