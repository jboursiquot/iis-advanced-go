package helloworld_test

import (
	"testing"

	"github.com/jboursiquot/iis-advanced-go/labs/helloworld"
)

func TestHelloWorld(t *testing.T) {
	got := helloworld.Hello()
	want := "Hello, world!"
	if got != want {
		t.Fatalf("Got %s but wanted %s", got, want)
	}
}

func TestHelloFromRSC(t *testing.T) {
	got := helloworld.HelloFromRSC()
	want := "Hello, world."
	if got != want {
		t.Fatalf("Got %s but wanted %s", got, want)
	}
}
