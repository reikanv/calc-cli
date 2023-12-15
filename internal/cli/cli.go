package cli

import (
	"flag"
	"fmt"

	"github.com/reikanv/calc-cli/internal/calc"
)

type XFlags struct {
	start float64
	end   float64
	delta float64
}

func ReadFlags() (xf XFlags, err error) {
	flag.Float64Var(&xf.start, "start", 0, "start value of range calculation")
	flag.Float64Var(&xf.end, "end", 0, "end value of range calculation")
	flag.Float64Var(&xf.delta, "deltaX", 0, "increments X")
	flag.Parse()

	if xf.delta > xf.end {
		err = fmt.Errorf("deltaX can't be more than end: deltaX: %v end: %v", xf.delta, xf.end)
		return
	}

	if xf.delta == 0 {
		err = fmt.Errorf("deltaX can't be zero")
		return
	}

	if xf.delta > 0 && xf.start >= xf.end {
		err = fmt.Errorf("invalid positive range: start %v end %v", xf.start, xf.end)
		return
	}

	if xf.delta < 0 && xf.start <= xf.end {
		err = fmt.Errorf("invalid negative range: end %v start %v", xf.end, xf.start)
		return
	}

	return xf, nil
}

func row(iter int, x float64) string {
	r, err := calc.Run(x)

	if err != nil {
		return fmt.Sprintf("%v %v ERROR", iter, x)
	}

	return fmt.Sprintf("%v %v %v", iter, x, r)
}

func Out(xf XFlags) {
	fmt.Println("# x result")
	i := 1

	for x := xf.start; x <= xf.end; x += xf.delta {
		fmt.Println(row(i, x))
		i += 1
	}
}
