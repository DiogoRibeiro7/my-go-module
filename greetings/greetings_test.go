package greetings

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Diogo")
	want := "Hello, Diogo!"
	if got != want {
		t.Fatalf("Hello(\"Diogo\")=%q; want %q", got, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	got := Hello("")
	want := "Hello, world!"
	if got != want {
		t.Fatalf("Hello(\"\")=%q; want %q", got, want)
	}
}
