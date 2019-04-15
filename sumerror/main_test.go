package main

import "testing"

func TestSum(t *testing.T) {
	i, _ := Sum("4", "5")
	if i != 9 {
		t.Errorf("The sum wasnt correct, got %d, want %d", i, 9)
	}
}

func TestSum_Error(t *testing.T) {
	_, err := Sum("faagf", "5")
	if err == nil {
		t.Error("There should be an error when converting invalid string", err)
	}
}
