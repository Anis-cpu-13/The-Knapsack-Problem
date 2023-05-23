package main_test

import (
	"testing"

	"./tools"
)

func BenchmarkGreedyAlgorithm(b *testing.B) {
	data, err := tools.LoadDataFromFile("data.json")
	if err != nil {
		b.Fatalf("Failed to load data: %v", err)
	}

	b.ReportAllocs()

	capacity := 80
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tools.SolveKnapsackWithGreedyAlgorithm(data, capacity)
	}
}

func BenchmarkDynamicProgramming(b *testing.B) {
	filename := "data.json"
	capacity := 80
	for n := 0; n < b.N; n++ {
		tools.SolveKnapsackWithDynamicProgramming(filename, capacity)
	}
}

func BenchmarkExhaustiveSearch(b *testing.B) {
	filename := "data.json"
	capacity := 80
	for n := 0; n < b.N; n++ {
		tools.SolveKnapsackWithExhaustiveSearch(filename, capacity)
	}
}
