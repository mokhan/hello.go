package main

import "testing"

func TestSum(test *testing.T) {
  numbers := [5]int{1, 2, 3, 4, 5}
  got := Sum(numbers)
  want := 15

  if got != want {
    test.Errorf("got %d want %d given, %v", got, want, numbers)
  }
}
