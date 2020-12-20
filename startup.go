package main

import (
	"log"
	"solve-problems-by-algorithms/solutions/baShuMa"
)

func main() {
	err := baShuMa.BSM()
	if err != nil {
		log.Fatal(err)
	}
}
