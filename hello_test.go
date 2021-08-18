package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("say hello", func(t *testing.T) {
		got := hello("world")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := hello("")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})
}
