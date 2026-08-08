package main

func indexOf[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

func indexOfTarges[T comparable](items []T, targets []T) int {
	for i, item := range items {
		for _, target := range targets {
			if item == target {
				return i
			}
		}
	}
	return -1
}
