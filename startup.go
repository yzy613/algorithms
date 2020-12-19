package main

import (
	"log"
	"solution-by-algorithm/solutions/baShuMa"
)

func main() {
	err := baShuMa.HRD()
	if err != nil {
		log.Fatal(err)
	}
}
