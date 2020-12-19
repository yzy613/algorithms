package main

import (
	"log"
	"solution-by-algorithm/solutions/baShuMa"
)

func main() {
	err := baShuMa.BSM()
	if err != nil {
		log.Fatal(err)
	}
}
