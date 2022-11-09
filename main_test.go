package main

import "testing"

func Test_Hello(t *testing.T) {
	want := "Hello Peter"

	got := hello("Peter")

	if want != got {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}
