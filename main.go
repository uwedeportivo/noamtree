package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"noamtree/noam"
)

var input = flag.String("input", "", "filename of file to convert")

func main() {
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	res, err := noam.ParseReader(*input, f)
	if err != nil {
		log.Fatal(err)
	}
	nodes := res.([]*noam.Node)
	fmt.Println(nodes[0])
}
