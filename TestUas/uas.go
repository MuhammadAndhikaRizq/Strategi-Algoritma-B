package main

import "fmt"

var (
	stops          = []int{10, 25, 30, 40, 50, 75, 80, 110, 130}
	targetDistance = 140
	solution       = []int{}
)

func findMinStops(currentDistance, currentIndex int) {
	if currentDistance > targetDistance {
		return
	}
	if currentDistance == targetDistance {
		fmt.Println(solution)
		return
	}
	for i := currentIndex; i < len(stops); i++ {
		solution = append(solution, stops[i])
		findMinStops(currentDistance+stops[i], i+1)
		solution = solution[:len(solution)-1]
	}
}

func main() {
	findMinStops(0, 0)
}
