package main

import "testing"

func TestSolve1(t *testing.T) {
	lines := []string{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"}
	got := solve1(lines)
	if got != 13 {
		t.Errorf("Result = %d; want 13", got)
	}
}

func TestSolve2A(t *testing.T) {
	lines := []string{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"}
	got := solve2(lines)
	if got != 1 {
		t.Errorf("Result = %d; want 1", got)
	}
}

func TestSolve2B(t *testing.T) {
	lines := []string{"R 5", "U 8", "L 8", "D 3", "R 17", "D 10", "L 25", "U 20"}
	got := solve2(lines)
	if got != 36 {
		t.Errorf("Result = %d; want 36", got)
	}
}
