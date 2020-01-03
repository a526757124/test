package main

import (
	"fmt"

	"github.com/a526757124/test/pipeline"
)

func main() {
	p := pipeline.InMenSort(
		pipeline.ArraySource(3, 2, 6, 7, 4,10,2,6,9,1,4,2,3,5,1,2,12,15,34,0,5,96,340,21))
	for v := range p {
		fmt.Print(v)
	}

}
