package main

import "testing"

func TestSolve1A(t *testing.T) {
	got := solve("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4)
	if got != 7 {
		t.Errorf("Result = %d; want 7", got)
	}
}

func TestSolve2A(t *testing.T) {
	got := solve("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)
	if got != 5 {
		t.Errorf("Result = %d; want 5", got)
	}
}

func TestSolve3A(t *testing.T) {
	got := solve("nppdvjthqldpwncqszvftbrmjlhg", 4)
	if got != 6 {
		t.Errorf("Result = %d; want 6", got)
	}
}

func TestSolve4A(t *testing.T) {
	got := solve("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)
	if got != 10 {
		t.Errorf("Result = %d; want 10", got)
	}
}

func TestSolve5A(t *testing.T) {
	got := solve("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4)
	if got != 11 {
		t.Errorf("Result = %d; want 11", got)
	}
}

func TestSolve1B(t *testing.T) {
	got := solve("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14)
	if got != 19 {
		t.Errorf("Result = %d; want 19", got)
	}
}

func TestSolve2B(t *testing.T) {
	got := solve("bvwbjplbgvbhsrlpgdmjqwftvncz", 14)
	if got != 23 {
		t.Errorf("Result = %d; want 23", got)
	}
}

func TestSolve3B(t *testing.T) {
	got := solve("nppdvjthqldpwncqszvftbrmjlhg", 14)
	if got != 23 {
		t.Errorf("Result = %d; want 23", got)
	}
}

func TestSolve4B(t *testing.T) {
	got := solve("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14)
	if got != 29 {
		t.Errorf("Result = %d; want 29", got)
	}
}

func TestSolve5B(t *testing.T) {
	got := solve("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14)
	if got != 26 {
		t.Errorf("Result = %d; want 26", got)
	}
}
