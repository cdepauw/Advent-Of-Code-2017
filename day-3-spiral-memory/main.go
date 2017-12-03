package main

import (
	"fmt"
	"math"
)

/*
You come across an experimental new kind of memory stored on an infinite two-dimensional grid.

Each square on the grid is allocated in a spiral pattern starting
at a location marked 1 and then counting up while spiraling outward.
For example, the first few squares are allocated like this:

17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...

While this is very space-efficient (no squares are skipped),
requested data must be carried back to square 1 (the location
of the only access port for this memory system) by programs
that can only move up, down, left, or right. They always take
the shortest path: the Manhattan Distance between the location
of the data and square 1.

For example:

- Data from square 1 is carried 0 steps, since it's at the access port.
- Data from square 12 is carried 3 steps, such as: down, left, left.
- Data from square 23 is carried only 2 steps: up twice.
- Data from square 1024 must be carried 31 steps.

Your puzzle input is 312051.
*/

/*
Ring max value:

1			Inner ring
+ 4*2 		Ring containing 2-9
+ 4*4		Ring containing 10-25
+ 4*6		Ring containing 26-49
+ 4*8		Ring containing 50-81
...
*/

func main() {
	// First we find the correct ring for this value
	goal := 312051

	// We start at the center, with a ring containing the value 1
	currentRingMaxValue := 1
	currentRingMinValue := 1
	currentRingNumber := 0 // We don't count the value 1 as being a ring

	// We start with ring size 0 and increment with 2 each step
	currentRingSize := 0

	// Search!
	for currentRingMaxValue < goal {
		currentRingNumber++
		currentRingSize += 2
		currentRingMinValue = currentRingMaxValue + 1
		currentRingMaxValue += 4 * currentRingSize
	}

	// Ring size 558 contains the value 312051.
	// This ring contains values between [310250, 312481]
	fmt.Printf("Ring #%v with size %v contains the value %v.\n", currentRingNumber, currentRingSize, goal)
	fmt.Printf("This ring contains values between [%v, %v].\n", currentRingMinValue, currentRingMaxValue)

	// The shortest Manhatten distance for this ring is 279 (currentRingNumber),
	// while the longest distance is 2*279 = 558.
	// We can find the distance of our goal, by finding its spot in the ring.

	// The number of values on the ring
	ringValuesCount := 4 * currentRingSize

	// The spot of our goal in this set of values
	goalValueLocation := goal - currentRingMinValue + 1

	// Our ring contains 2232 values and our goal is value #1802
	fmt.Printf("Our ring contains %v values and our goal is value #%v\n", ringValuesCount, goalValueLocation)

	// Divide into 4 quadrants, one for each side, and see how far
	// our goal is from the center value of that quadrant.
	ringSideValuesCount := ringValuesCount / 4

	quadrant := 1
	quadrantMinValue := 0
	quadrantMaxValue := ringSideValuesCount
	for goalValueLocation > quadrantMaxValue {
		quadrant++
		quadrantMinValue = quadrantMaxValue
		quadrantMaxValue += ringSideValuesCount
	}

	// Our goal is located in quadrant 4, with value locations 1674 to 2232.
	fmt.Printf("Our goal is located in quadrant %v, with value locations %v to %v.\n", quadrant, quadrantMinValue, quadrantMaxValue)

	quadrantCenterValue := quadrantMinValue + (quadrantMaxValue-quadrantMinValue)/2
	goalCenterOffset := int(math.Abs(float64(goalValueLocation - quadrantCenterValue)))

	// Our quadrant center value is 1953, meaning our goal is 151 steps away from the center value.
	fmt.Printf("Our quadrant center value is %v, meaning our goal is %v steps away from the center value.\n", quadrantCenterValue, goalCenterOffset)

	goalManhattenDistance := currentRingNumber + goalCenterOffset

	// The resulting Manhatten distance for our goal is 430
	fmt.Printf("The resulting Manhatten distance for our goal is %v", goalManhattenDistance)

	// One pretty blank line to wrap this up
	fmt.Println()
}
