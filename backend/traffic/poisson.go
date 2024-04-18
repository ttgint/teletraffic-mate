package traffic

import "math"

type Poisson struct {
	Measurements Measurement
}

func (p Poisson) Calc_gos(N int, A float64) float64 {
	temp := 0.0
	for t := 0; t < N; t++ {
		temp += math.Pow(A, float64(t)) / FloatFactorial(float64(t))
	}
	p.Measurements.grade_of_service = 1 - (math.Exp(-A) * temp)
	return Round(p.Measurements.grade_of_service, 3)
}

func (p Poisson) Calc_traffic(num_of_channels int, GoS float64) float64 {
	erlangC := new(ErlangC)
	trunks := float64(num_of_channels)
	if num_of_channels < 1 || GoS < 0 {
		return 0
	}

	var maxI, B, incr float64
	maxI = trunks
	B = erlangC.ErlangC(num_of_channels, maxI)

	for B < GoS {
		maxI *= 2
		B = erlangC.ErlangC(num_of_channels, maxI)
	}

	incr = 1
	for incr <= maxI/100 {
		incr *= 10
	}
	p.Measurements.traffic = LoopingTraffic(trunks, GoS, incr, maxI, 0, 1)
	return Round(p.Measurements.traffic, 3)
}

func (p Poisson) Calc_num_channels(A float64, GS float64) float64 {
	var NN, err, GoS []float64

	NN = make([]float64, 100000)
	err = make([]float64, 100000)
	GoS = make([]float64, 100000)

	NN[0] = 1
	err[0] = 10000

	for tt := 0; tt < 10000; tt++ {
		N := NN[tt]
		ara := 0.0

		for t := 0; t < int(N); t++ {
			ara += math.Pow(A, float64(t)) / FloatFactorial(float64(t))
		}

		GoS[tt] = 1 - (math.Exp(-A) * ara)
		err[tt+1] = math.Abs(GoS[tt] - GS)

		if err[tt+1]-err[tt] > 0 {
			N = NN[tt] - 1
			p.Measurements.traffic = N
			return Round(p.Measurements.traffic, 3)
		}

		NN[tt+1] = N + 1
	}

	return 0
}
