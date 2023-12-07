package boat

import (
	"fmt"
	"math"
)

func calculateX12ToWin(t, recordDistance float64) (int, int) {

	x1 := t/2 + math.Sqrt((math.Pow(t, 2)/4)-recordDistance)
	x2 := t/2 - math.Sqrt((math.Pow(t, 2)/4)-recordDistance)

	// log the results
	fmt.Printf("x1: %f, x2: %f\n", x1, x2)

	// first of all order that x1 is the lesser number
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	// if x1 is a whole number we need to increase it by 1
	if x1 == math.Floor(x1) {
		x1++
	}

	// if x2 is a whole number we need to decrease it by 1
	if x2 == math.Floor(x2) {
		x2--
	}

	// x1 always needs to be rounded up
	x1 = math.Ceil(x1)

	// x2 always needs to be rounded down
	x2 = math.Floor(x2)

	// log the results
	fmt.Printf("x1: %f, x2: %f\n", x1, x2)

	return int(math.RoundToEven(x1)), int(math.RoundToEven(x2))

}

func getY(x, t float64) int {
	return int((-1*math.Pow(x, 2) + x*t))
}

func calcNumberOfWinning(t, recordDistance float64) int {

	x1, x2 := calculateX12ToWin(t, recordDistance)
	return x2 - x1 + 1

}
