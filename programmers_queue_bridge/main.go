package main

import (
	"fmt"
)

var (
	bridge_length = 2
	weight        = 10
	truck_weights = []int{7, 4, 5, 6}
)

func solution(bridge_length int, weight int, truck_weights []int) int {
	type truck struct {
		bridge_length int
		truck_weight  int
	}

	goingTrucks := make([]truck, 0)

	bridge_weight := 0
	result := 1
	for {
		if len(truck_weights) != 0 && bridge_weight+truck_weights[0] <= weight {
			nowTruck := truck_weights[0]
			truck_weights = truck_weights[1:]
			Truck := truck{
				bridge_length: bridge_length,
				truck_weight:  nowTruck,
			}

			bridge_weight += Truck.truck_weight
			goingTrucks = append(goingTrucks, Truck)
		}

		isMoving := 0
		truckLen := len(goingTrucks)
		isComplete := false
		for i := 0; i < truckLen; i++ {
			isMoving++
			goingTrucks[i].bridge_length--
			if goingTrucks[i].bridge_length == 0 {
				bridge_weight -= goingTrucks[i].truck_weight
				isComplete = true
			}
		}
		if isComplete == true {
			goingTrucks = goingTrucks[1:]
		}

		if isMoving == 0 {
			break
		} else {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(solution(bridge_length, weight, truck_weights))
}
