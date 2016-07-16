package main

import (
	"fmt"
    "github.com/goris/algorithm-design/reader"
    "github.com/goris/algorithm-design/inversions"
	"log"
)

func main() {
	lines, err := reader.ReadInts("data/integersInversions.txt")
	if err != nil {
		log.Fatalf("readInts: %s", err)
	}
    inversions := inversions.CountInversionsO2(lines)
    fmt.Println("# of inversions: ", inversions);

}
