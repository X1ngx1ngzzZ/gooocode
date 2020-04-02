package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if want != got {
		t.Error("ADD 1 and 2 result isn't 3")
	}
}
