package traffic

import (
	"math"
)

type Measurement struct {
	num_of_channels  int
	traffic          float64
	grade_of_service float64 //(percantage)
}

func Factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return (num * Factorial(num-1))
}
func FloatFactorial(f float64) float64 {
	if f <= 1.0 {
		return 1.0
	}
	return (f * FloatFactorial(f-1))
}

// round rounds a float64 to a specified number of decimal places
func Round(x float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	xShifted := x * shift
	rounded := math.Round(xShifted)
	result := rounded / shift
	return result
}

// Ensuring the value is between [max, min]
// Otherwise, floor or ceil the value
func MinMax(val float64, min, max int) float64 {
	minmax := val
	if val < float64(min) {
		minmax = float64(min)
		return minmax
	}
	if val > float64(max) {
		minmax = float64(max)
		return minmax
	}
	return minmax
}

// LoopingTraffic performs an iterative search to find the minimum traffic intensity
// for a target Grade of Service (GoS) using Erlang B or Erlang C models.
//   Increment: Initial intensity increment.
//   MaxIntensity: Upper limit for traffic intensity.
//   MinIntensity: Lower limit for traffic intensity.
// Returns the minimum traffic intensity that achieves the specified GoS.
// The algorithm uses a binary search-like approach, adjusting intensity until
// GoS matches the target or maximum accuracy is reached within a defined number of loops.
func LoopingTraffic(Truncks, GoS, Increment, MaxInetensity, MinInetensity float64, erlangType int) float64 {
	erlangb := new(ErlangB)
	erlangc := new(ErlangC)
	var incr, B, minI, intensity, counter float64
	MaxLoops := 100
	MaxAccuracy := 0.00001

	minI = MinInetensity
	if erlangType == 0 {
		B = erlangb.ErlangB(int(Truncks), minI)
	} else if erlangType == 1 {
		B = erlangc.ErlangC(int(Truncks), minI)
	}

	if B == GoS {
		Traffic := minI
		return Traffic
	}
	incr = Increment
	intensity = minI
	counter = 0

	for (incr >= MaxAccuracy) && (counter < float64(MaxLoops)) {
		if erlangType == 0 {
			B = erlangb.ErlangB(int(Truncks), intensity)
		} else if erlangType == 1 {
			B = erlangc.ErlangC(int(Truncks), intensity)
		}
		if B > GoS {
			incr /= 10
			intensity = minI
		}
		minI = intensity
		intensity += incr
		counter++
	}
	return minI
}
