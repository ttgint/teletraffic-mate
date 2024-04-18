package traffic

import "math"

type ErlangC struct {
	Measurements Measurement
}

// Calculating ErlangB (GoS) in an iterative approach
func (e ErlangC) ErlangC(num_of_channels int, traffic float64) float64 {
	erlangb := new(ErlangB)
	var b, c float64
	if (num_of_channels < 0) || (traffic < 0) {
		return 0
	}
	b = erlangb.ErlangB(num_of_channels, traffic)
	c = b / (((traffic / float64(num_of_channels)) * b) + (1 - (traffic / float64(num_of_channels))))

	ErlangC := MinMax(c, 0, 1)
	return Round(ErlangC, 3)
}

func (e ErlangC) Calc_num_channels(A float64, GS float64) float64 {
	const maxIterations = 10000

	NN := make([]float64, maxIterations)
	err := make([]float64, maxIterations)
	GoS := make([]float64, maxIterations)

	NN[0] = math.Trunc(A)
	err[0] = 10000

	for tt := 0; tt < maxIterations; tt++ {
		N := NN[tt]
		rslt := 1.0
		kats := 0.0
		E := 0.0

		for t := 1; t < int(N); t++ {
			rslt += math.Pow(A, float64(t)) / FloatFactorial(float64(t))
		}

		kats = (N * math.Pow(A, N)) / ((N - A) * FloatFactorial(N))
		E = 1 / (kats + rslt)
		GoS[tt] = kats * E
		err[tt+1] = math.Abs(GoS[tt] - GS)

		if err[tt+1]-err[tt] > 0 {
			N = NN[tt] - 1
			return N
		} else {
			NN[tt+1] = N + 1
		}
	}
	return -1
}

func (e ErlangC) Calc_traffic(num_of_channels int, GoS float64) float64 {
	var b, incr, trunck, maxI float64
	trunck = float64(num_of_channels)
	if trunck < 1.0 || GoS < 0 {
		return 0
	}

	maxI = trunck
	b = e.ErlangC(num_of_channels, maxI)
	for b < GoS {
		maxI *= 2
		b = e.ErlangC(num_of_channels, maxI)
	}
	incr = 1
	for incr <= ((maxI) / 100) {
		incr *= 10
	}
	e.Measurements.traffic = LoopingTraffic(trunck, GoS, incr, maxI, 0, 1)
	return Round(e.Measurements.traffic, 3)
}

func (e ErlangC) Calc_gos(N int, A float64) float64 {
	rslt := 1.0
	var t int
	for t = 1; t <= N-1; t++ {
		rslt += math.Pow(A, float64(t)) / float64(Factorial(t))
	}

	kats := (float64(N) * math.Pow(A, float64(N))) / ((float64(N) - A) * float64(Factorial(t)))
	E := 1 / (kats + rslt)
	e.Measurements.grade_of_service = (kats * E)
	return Round(e.Measurements.grade_of_service, 3)
}
