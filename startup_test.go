package main

import (
	"solution-by-algorithm/solutions/baShuMa"
	"testing"
)

func TestStartup(t *testing.T) {
	err := baShuMa.BSM()
	if err != nil {
		t.Error(err)
	}
}
