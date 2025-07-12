package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Geir")
	want := "Hello, Geir!"
	if got != want {
		t.Errorf("Hello() = %q; want %q", got, want)
	}
}
