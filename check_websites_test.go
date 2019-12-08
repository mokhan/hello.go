package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://www.example.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com",
		"https://www.example.com",
	}

	want := map[string]bool{
		"https://www.google.com":  true,
		"https://www.example.com": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v got %v", want, got)
	}
}
