package main

import (
	"fmt"
	"math"
)

func mostRepeatedElement(arr []string) string {
	counts := make(map[string]int)

	for _, item := range arr {
		counts[item]++
	}

	maxCount := math.MinInt
	var mostRepeated string

	for item, count := range counts {
		if count > maxCount {
			maxCount = count
			mostRepeated = item
		}
	}

	return mostRepeated
}

func main() {
	inputData := []string{"apple", "pie", "apple", "red", "red", "red"}
	output := mostRepeatedElement(inputData)
	fmt.Println(output)
}
