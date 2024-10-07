package main

import (
	"fmt"

	"noamtree/noam"
)

const (
	example = `[ CP [C that] [TP [NP They] [VP [V know] [CP [C that] [TP [NP she]
               [VP [V knows] [CP [ C that] [TP [NP you] [VP [V know] [CP [C that]
               [TP [NP I] [VP [V hate] [NP war] ] ]]]]]]]]]]]`
	example2 = `[CP foo]`
	example3 = `[CP [C that]]`
	example4 = `[CP [C that] [C this]]`
)

func main() {
	res, err := noam.Parse("example.noam", []byte(example))
	if err != nil {
		panic(err)
	}
	node := res.(*noam.Node)
	fmt.Println(node)
}
