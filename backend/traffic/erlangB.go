package traffic

import (
	"math"
)

type ErlangB struct {
	Measurements Measurement
}

// Calculating ErlangB (GoS) in an iterative approach
func (e ErlangB) ErlangB(num_of_channels int, traffic float64) float64 {
	var maxI int
	var val, last, B float64

	if (num_of_channels < 0) || (traffic < 0) {
		return 0
	}

	maxI = num_of_channels
	val = traffic
	last = 1
	for i := 1; i <= maxI; i++ {
		B = (val * last) / (float64(i) + (val * last))
		last = B
	}

	ErlangB := MinMax(B, 0, 1)
	return ErlangB
}

// ErlangB Based Traffic Calculations
func (e ErlangB) Calc_traffic(num_of_channels int, GoS float64) float64 {
	var b, incr, maxI, trunck float64
	trunck = float64(num_of_channels)
	if num_of_channels < 1.0 || GoS < 0.0 {
		return 0
	}

	maxI = trunck
	b = e.ErlangB(num_of_channels, maxI)
	for b < GoS {
		maxI = maxI * 2
		b = e.ErlangB(num_of_channels, maxI)
	}
	incr = 1.0
	for incr <= (maxI / 100) {
		incr *= 10
	}
	e.Measurements.traffic = LoopingTraffic(trunck, GoS, incr, maxI, 0, 0)
	return Round(e.Measurements.traffic, 3)
}

func (e ErlangB) Calc_num_channels(traffic, GoS float64) int {
	if GoS < 0 || traffic < 0 {
		return 0
	}

	val := traffic
	last := 1.0
	B := 1.0
	counter := 0.0

	for (B > GoS) && (B > 0.001) {
		counter++
		B = (val * last) / (counter + (val * last))
		last = B
	}

	e.Measurements.num_of_channels = int(math.Round(counter))
	return e.Measurements.num_of_channels
}

// Calculate the grade of serivce using ErlangB equation
func (e ErlangB) Calc_gos(num_of_channels int, traffic float64) float64 {
	rslt := 1.0
	E := 1.0

	for t := num_of_channels; t >= 1; t-- {
		rslt *= float64(t) / traffic
		E += rslt
	}
	e.Measurements.grade_of_service = Round(1/E, 4)
	return e.Measurements.grade_of_service
}
