package main

/*
--- Part Two ---

As a stress test on the system, the programs here clear the grid
and then store the value 1 in square 1. Then, in the same allocation
order as shown above, they store the sum of the values in all adjacent
squares, including diagonals.

So, the first few squares' values are chosen as follows:

Square 1 starts with the value 1.
Square 2 has only one adjacent filled square (with value 1), so it also stores 1.
Square 3 has both of the above squares as neighbors and stores the sum of their values, 2.
Square 4 has all three of the aforementioned squares as neighbors and stores the sum of their values, 4.
Square 5 only has the first and fourth squares as neighbors, so it gets the value 5.
Once a square is written, its value does not change. Therefore, the first few squares would receive the following values:

147  142  133  122   59
304    5    4    2   57
330   10    1    1   54
351   11   23   25   26
362  747  806  880  931 957
What is the first value written that is larger than your puzzle input?

*/

// SpiralStore calculates the step 2 result
func SpiralStore(goal int) int {
	start := 10
	current := 0
	for current < goal {
		current = SpiralSummedValue(start)
		start++
	}

	return current
}

// SpiralSummedValue calculates the "summed" value of a
// given number (goal) as described above.
// PreviousGoals expects the preceding goals to be given,
// and as such len(previousGoals) == goal-1.
func SpiralSummedValue(goal int) int {
	hardCodeded := []int{0, 1, 1, 2, 4, 5, 10, 11, 23, 25}
	if goal < len(hardCodeded) {
		return hardCodeded[goal]
	}

	// Calculate the Manhatten distance
	manhattenDistance := ManhattenDistance(goal)

	// Calculate ring information
	// currentRingMaxValue, currentRingMinValue, currentRingNumber, currentRingSideSize := CalculateRingInformation(goal)
	_, currentRingMinValue, currentRingNumber, currentRingSideSize := CalculateRingInformation(goal)
	previousRingMaxValue := currentRingMinValue - 1
	previousRingMinValue := currentRingMinValue - 4*(currentRingSideSize-2)
	if currentRingNumber == 1 {
		previousRingMinValue = 1
		previousRingMaxValue = 1
	}

	// The number of values on the ring
	ringValuesCount := 4 * currentRingSideSize

	// Divide into 4 quadrants, one for each side, and see how far
	// our goal is from the center value of that quadrant.
	ringSideValuesCount := ringValuesCount / 4

	// Calculate the quadrant.
	// Quadrant 1 = East
	// Quadrant 2 = North
	// Quadrant 3 = West
	// Quadrant 4 = South
	quadrant := 1
	quadrantMinValue := currentRingMinValue
	quadrantMaxValue := quadrantMinValue + ringSideValuesCount - 1
	for goal > quadrantMaxValue {
		quadrant++
		quadrantMinValue = quadrantMaxValue + 1
		quadrantMaxValue = quadrantMinValue + ringSideValuesCount - 1
	}

	// The inner neighbour is part of finding all the required neighbour values
	innerNeighbourIndex := CalculateInnerNeighourIndex(goal, quadrant, currentRingSideSize, currentRingMinValue, currentRingNumber, manhattenDistance)

	// Calculate the final result
	//fmt.Println("Goal:", goal, ", Quad min:", quadrantMinValue, ", Quad max:", quadrantMaxValue, ", Prev MinVal:", previousRingMinValue, ", Prev MaxVal:", previousRingMaxValue, ", Inner idx:", innerNeighbourIndex)
	result := SpiralSummedValue(innerNeighbourIndex)
	switch quadrant {
	case 1:
		if goal == quadrantMinValue {
			result += SpiralSummedValue(previousRingMinValue)
		} else if goal == quadrantMinValue+1 {
			result += SpiralSummedValue(previousRingMinValue + 1)
			result += SpiralSummedValue(previousRingMaxValue)
			result += SpiralSummedValue(previousRingMaxValue + 1)
		} else if goal == quadrantMaxValue {
			result += SpiralSummedValue(goal - 1)
		} else if goal == quadrantMaxValue-1 {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(innerNeighbourIndex - 1)
		} else {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(innerNeighbourIndex - 1)
			result += SpiralSummedValue(innerNeighbourIndex + 1)
		}
	case 2:
		if goal == quadrantMaxValue {
			result += SpiralSummedValue(goal - 1)
		} else if goal == quadrantMaxValue-1 {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(innerNeighbourIndex - 1)
		} else if goal == quadrantMinValue {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(goal - 2)
			result += SpiralSummedValue(innerNeighbourIndex + 1)
		} else {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(innerNeighbourIndex - 1)
			result += SpiralSummedValue(innerNeighbourIndex + 1)
		}
	case 3:
		if goal == quadrantMaxValue {
			result += SpiralSummedValue(goal - 1)
		} else if goal == quadrantMaxValue-1 {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(innerNeighbourIndex - 1)
		} else if goal == quadrantMinValue {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(goal - 2)
			result += SpiralSummedValue(innerNeighbourIndex + 1)
		} else {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(innerNeighbourIndex - 1)
			result += SpiralSummedValue(innerNeighbourIndex + 1)
		}
	case 4:
		if goal == quadrantMaxValue-1 {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(innerNeighbourIndex - 1)
			result += SpiralSummedValue(innerNeighbourIndex + 1)
		} else if goal == quadrantMaxValue {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(innerNeighbourIndex + 1)
		} else if goal == quadrantMinValue {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(goal - 2)
			result += SpiralSummedValue(innerNeighbourIndex + 1)
		} else {
			result += SpiralSummedValue(goal - 1)
			result += SpiralSummedValue(innerNeighbourIndex - 1)
			result += SpiralSummedValue(innerNeighbourIndex + 1)
		}
	}

	return result
}

// CalculateInnerNeighourIndex calculates the inner neighbour
// of the given value (goal). This depends on quadrant,
// ring number and ring size.
func CalculateInnerNeighourIndex(goal, quadrant, ringSideSize, minRingValue, ringNumber, manhattenDistance int) int {
	// The first actual ring always reverts to value 1 as their inner neighbour.
	if goal < 10 {
		return 1
	}

	// The first value of a ring comes directly after its inner neighbour.
	if goal == minRingValue {
		return goal - 1
	}

	// If the goal is at the max Manhatten distance, its inner neighbour
	// is located via a diagonal. This requires an index offset of 1.
	manhattenDistanceOffset := 0
	if manhattenDistance == ringNumber+int(ringSideSize/2) {
		manhattenDistanceOffset = -1
	}

	// ringOffset is the number of values in the previous ring +
	// the number of the previous ring.
	ringOffset := (ringSideSize-2)*4 + (ringNumber - 1)

	// The total offset is the ringOffset +
	totalOffset := ringOffset + ((quadrant - 1) * 2) - (ringNumber - 2)

	return goal - totalOffset + manhattenDistanceOffset
}
