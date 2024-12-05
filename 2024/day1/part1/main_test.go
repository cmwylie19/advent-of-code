package main

import (
	"testing"
)

func TestSortArray(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"already sorted", []int{1, 2, 3, 4, 5}, expected},
		{"reverse order", []int{5, 4, 3, 2, 1}, expected},
		{"random order", []int{3, 1, 4, 5, 2}, expected},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArray(tt.input)
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("sortArray() = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}
