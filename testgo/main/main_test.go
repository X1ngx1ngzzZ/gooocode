package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Error("ADD 1 and 2 result isn't 3")
	}
}
