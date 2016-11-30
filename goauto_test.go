package goauto

import (
	"testing"
)

func TestAddTwo(t *testing.T) {
	got := AddTwo(3, 7)
	want := 10
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAddThree(t *testing.T) {
	got := AddThree(3, 7, 11)
	want := 21
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAddFour(t *testing.T) {
	got := AddFour(3, 7, 11, 13)
	want := 34
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAddFive(t *testing.T) {
	got := AddFive(3, 7, 11, 13, 17)
	want := 51
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}