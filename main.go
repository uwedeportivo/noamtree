package main

import (
	"flag"
	"fmt"
	"log"

	"noamtree/noam"
)

var input = flag.String("input", "", "filename of file to convert")

func main() {
	flag.Parse()

	res, err := noam.ParseFile(*input)
	if err != nil {
		log.Fatal(err)
	}
	nodes := res.([]*noam.Node)
	fmt.Println(nodes[0])
}
