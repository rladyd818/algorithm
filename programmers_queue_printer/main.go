package main

import (
	"fmt"
)

var (
	// priorities = []int{2, 1, 3, 2}
	priorities = []int{1, 1, 9, 1, 1, 1}
	// location = 2
	location = 0
)

func solution(priorities []int, location int) int {
	result := 0
	docLocation := location
	for {
		value := priorities[0]
		priorities = priorities[1:]
		isPrint := true
		for i := 0; i < len(priorities); i++ {
			if value < priorities[i] {
				isPrint = false
				break
			}
		}

		if isPrint {
			result++
			if docLocation == 0 {
				break
			}
		} else {
			priorities = append(priorities, value)
		}

		docLocation--
		if docLocation == -1 {
			docLocation = len(priorities) - 1
		}
	}

	return result
}

func main() {
	fmt.Println(solution(priorities, location))
}
