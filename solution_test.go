package square

import (
	"math"
	"testing"
)

type testCase struct {
	sideLen        float64
	sidesNum       ShapeSides
	expectedResult float64
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

var testCases = []testCase{
	{2, SidesSquare, 4},
	{4, SidesSquare, 16},
}

func TestCalcSquare(t *testing.T) {
	for _, tc := range testCases {
		result := CalcSquare(tc.sideLen, tc.sidesNum)
		if !almostEqual(result, tc.expectedResult) {
			t.Errorf("TestCalcSquare: failed case sidelen=%v, sidesNum=%q, expected=%v. Actual=%v", tc.sideLen, tc.sidesNum, tc.expectedResult, result)
		}
	}
}
