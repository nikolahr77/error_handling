package main

import "testing"

func TestSum(t *testing.T) {
	i, err := Sum("4", "5")
	if i != 9 {
		t.Errorf("The sum wasnt correct, got %d, want %d", i, 9)
	}
	if err != nil{
		t.Error("Unexpected error:", err)
	}
}

func TestSum_PutNonIntegerFirstParameter(t *testing.T) {
	_, err := Sum("faagf", "5")
	if err == nil {
		t.Error("There should be an error when converting invalid string", err)
	}
}

func TestSum_PutNonIntegerSecondParameter(t *testing.T) {
	_, err := Sum("4","hello")
	if err == nil{
		t.Errorf("There should be an error when second integer is invalid")
	}
}