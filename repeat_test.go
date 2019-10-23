package main

import "testing"

func TestRepeat(t *testing.T) {
  t.Run("default", func(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
      t.Errorf("expected %q but got %q", expected, repeated)
    }
  })

  t.Run("custom interval", func(t *testing.T) {
    repeated := Repeat("a", 10)
    expected := "aaaaaaaaaa"

    if repeated != expected {
      t.Errorf("expected %q but got %q", expected, repeated)
    }
  })
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a")
  }
}
