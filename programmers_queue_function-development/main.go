package main

import (
	"fmt"
)

var (
	progresses = []int{93, 30, 55}
	speeds     = []int{1, 30, 5}
)

func solution(progresses []int, speeds []int) []int {
	complete := make([]bool, len(speeds))
	result := make([]int, 0)
	// 100%가 되는 날 배포, slice의 index순서대로 배포되어야함

	for {
		for i := range speeds {
			progresses[i] += speeds[i]
			if progresses[i] >= 100 {
				complete[i] = true
			}
		}

		if len(progresses) != 0 {
			if complete[0] {
				slicePtr := 0
				for i := range complete {
					if complete[i] {
						slicePtr = i + 1
					} else {
						break
					}
				}
				result = append(result, len(progresses[:slicePtr]))
				progresses = progresses[slicePtr:]
				speeds = speeds[slicePtr:]
				complete = complete[slicePtr:]
			}
		} else {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(solution(progresses, speeds))
}
