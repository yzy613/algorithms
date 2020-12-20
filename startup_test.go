package main

import (
	"solve-problems-by-algorithms/solutions/baShuMa"
	"testing"
)

func TestStartup(t *testing.T) {
	err := baShuMa.BSM()
	if err != nil {
		t.Error(err)
	}
}
