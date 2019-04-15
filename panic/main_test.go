package main

import "testing"

func TestSum(t *testing.T) {
	i := Sum(5,10)
	if i!=15 {
		t.Errorf("The sum wasnt correct, got %d, want %d",i,15)
	}
}
