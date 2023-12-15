package calc

import (
	"fmt"
	"math"
)

const num = 97

func calcEven(x float64) float64 {
	p1 := math.Log((1 - num) / math.Sin(x+num))
	p2 := math.Abs(1 / math.Tan(x) / num)
	return math.Max(p1, p2)
}

func calcOdd(x float64) float64 {
	p1 := math.Log((1 - num) / math.Cos(x-num))
	p2 := math.Tan(x) / num
	return math.Min(p1, p2)
}

func Run(x float64) (r float64, err error) {
	if math.Mod(x, 2) == 0 {
		r = calcEven(x)
	} else {
		r = calcOdd(x)
	}

	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, fmt.Errorf("ERROR")
	}

	return
}
