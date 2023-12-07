package boat

import (
	"fmt"
	"testing"
)

func TestCalculateX12ToWin(t *testing.T) {

	tests := []struct {
		t              float64
		recordDistance float64
		expectedX1     int
		expectedX2     int
	}{
		{
			t:              7,
			recordDistance: 9,
			expectedX1:     2,
			expectedX2:     5,
		},
	}

	for _, test := range tests {
		x1, x2 := calculateX12ToWin(test.t, test.recordDistance)

		// the order of the return values is not guaranteed
		if (x1 != test.expectedX1 && x1 != test.expectedX2) ||

			(x2 != test.expectedX1 && x2 != test.expectedX2) {
			t.Errorf("Test Case: t: %f, recordDistance: %f\nExpected: %d, %d\nGot: %d, %d", test.t, test.recordDistance, test.expectedX1, test.expectedX2, x1, x2)

		}

	}
}

func TestCalcNumberOfWinning(t *testing.T) {

	tests := []struct {
		t              float64
		recordDistance float64
		expectedResult int
	}{
		{
			t:              7,
			recordDistance: 9,
			expectedResult: 4,
		},
		{
			t:              15,
			recordDistance: 40,
			expectedResult: 8,
		},
		{
			t:              30,
			recordDistance: 200,
			expectedResult: 9,
		},
	}

	for _, test := range tests {
		result := calcNumberOfWinning(test.t, test.recordDistance)

		if result != test.expectedResult {
			t.Errorf("Test Case: t: %f, recordDistance: %f\nExpected: %d\nGot: %d", test.t, test.recordDistance, test.expectedResult, result)
		}
	}

}

// create new test with

func TestExample1(t *testing.T) {

	// Time:        54     70     82     75
	// Distance:   239   1142   1295   1253

	tests := []struct {
		t              float64
		recordDistance float64
		expectedResult int
	}{
		{
			t:              54,
			recordDistance: 239,
			expectedResult: 45,
		},
		{
			t:              70,
			recordDistance: 1142,
			expectedResult: 19,
		},
		{
			t:              82,
			recordDistance: 1295,
			expectedResult: 39,
		},
		{
			t:              75,
			recordDistance: 1253,
			expectedResult: 24,
		},
	}

	multi := 1
	for _, test := range tests {

		result := calcNumberOfWinning(test.t, test.recordDistance)
		// log the results
		fmt.Printf("t: %f, recordDistance: %f\n", test.t, test.recordDistance)

		multi = multi * result

		if result != test.expectedResult {
			t.Errorf("Test Case: t: %f, recordDistance: %f\nExpected: %d\nGot: %d", test.t, test.recordDistance, test.expectedResult, result)
		}

		// log multi
		fmt.Printf("multi: %d\n", multi)

	}
}

func TestExample2(t *testing.T) {

	// Time:        54     70     82     75
	// Distance:   239   1142   1295   1253

	tests := []struct {
		t              float64
		recordDistance float64
		expectedResult int
	}{
		{
			t:              71530,
			recordDistance: 940200,
			expectedResult: 71503,
		},
	}

	multi := 1
	for _, test := range tests {

		result := calcNumberOfWinning(test.t, test.recordDistance)
		// log the results
		fmt.Printf("t: %f, recordDistance: %f\n", test.t, test.recordDistance)

		multi = multi * result

		if result != test.expectedResult {
			t.Errorf("Test Case: t: %f, recordDistance: %f\nExpected: %d\nGot: %d", test.t, test.recordDistance, test.expectedResult, result)
		}

		// log multi
		fmt.Printf("multi: %d\n", multi)

	}
}

func TestData2(t *testing.T) {

	// Time:        54     70     82     75
	// Distance:   239   1142   1295   1253

	tests := []struct {
		t              float64
		recordDistance float64
		expectedResult int
	}{
		{
			t:              54708275,
			recordDistance: 239114212951253,
			expectedResult: 45128024,
		},
	}

	multi := 1
	for _, test := range tests {

		result := calcNumberOfWinning(test.t, test.recordDistance)
		// log the results
		fmt.Printf("t: %f, recordDistance: %f\n", test.t, test.recordDistance)

		multi = multi * result

		if result != test.expectedResult {
			t.Errorf("Test Case: t: %f, recordDistance: %f\nExpected: %d\nGot: %d", test.t, test.recordDistance, test.expectedResult, result)
		}

		// log multi
		fmt.Printf("multi: %d\n", multi)

	}
}
