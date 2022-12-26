package main

import (
	"strconv"
	"testing"
)

func TestSnafuToDec(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"1=-0-2", 1747},
		{"12111", 906},
		{"2=0=", 198},
		{"21", 11},
		{"2=01", 201},
		{"111", 31},
		{"20012", 1257},
		{"112", 32},
		{"1=-1=", 353},
		{"1-12", 107},
		{"12", 7},
		{"1=", 3},
		{"122", 37},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ans := snafuToDec(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestDecToSnafu(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{1747, "1=-0-2"},
		{906, "12111"},
		{198, "2=0="},
		{11, "21"},
		{201, "2=01"},
		{31, "111"},
		{1257, "20012"},
		{32, "112"},
		{353, "1=-1="},
		{107, "1-12"},
		{7, "12"},
		{3, "1="},
		{37, "122"},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			ans := decToSnafu(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
