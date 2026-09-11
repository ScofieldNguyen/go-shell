package main

import "testing"

func TestIndexOfTargets(t *testing.T) {
	tests := []struct {
		name    string
		items   []string
		targets []string
		want    int
	}{
		{"first item matches", []string{"echo", "something", "2>", "output.txt"}, []string{">>", "1>>"}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := indexOfTarges(tt.items, tt.targets)
			if got != tt.want {
				t.Errorf("indexOfTarges(%v, %v) = %d, want %d", tt.items, tt.targets, got, tt.want)
			}
		})
	}
}
