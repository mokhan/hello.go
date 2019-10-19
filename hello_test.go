package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("mo")
  want := "Hello, mo"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
