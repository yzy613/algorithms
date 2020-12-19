package main

import (
	"solution-by-algorithm/solutions/baShuMa"
	"testing"
)

func TestStartup(t *testing.T) {
	err := baShuMa.HRD()
	if err != nil {
		t.Error(err)
	}
}
