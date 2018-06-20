package main

import (
	"fmt"

	"github.com/a526757124/test/pipeline"
)

func main() {
	p := pipeline.InMenSort(
		pipeline.ArraySource(3, 2, 6, 7, 4))
	for v := range p {
		fmt.Println(v)
	}

}
