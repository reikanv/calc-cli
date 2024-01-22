package cli

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/reikanv/calc-cli/internal/calc"
	"github.com/reikanv/calc-cli/pkg/strpad"
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

	if (xf.delta > xf.end && xf.end > 0) || (xf.delta < xf.end && xf.end < 0) {
		err = fmt.Errorf("invalid deltaX. deltaX: %v end: %v", xf.delta, xf.end)
		return
	}

	if xf.delta == 0 {
		err = fmt.Errorf("deltaX can't be zero")
		return
	}

	if xf.delta > 0 && xf.start >= xf.end {
		err = fmt.Errorf("invalid positive range. start %v end %v", xf.start, xf.end)
		return
	}

	if xf.delta < 0 && xf.start <= xf.end {
		err = fmt.Errorf("invalid negative range. end %v start %v", xf.end, xf.start)
		return
	}

	return xf, nil
}

func row(iter int, x float64, pad int) string {
	r, err := calc.Run(x)
	paddedIter := strpad.Right(fmt.Sprintf("%v", iter), pad)
	paddedX := strpad.Right(fmt.Sprintf("%.6f", x), pad)

	if err != nil {
		return fmt.Sprintf("%v | %v | ERROR", paddedIter, paddedX)
	}

	paddedR := strpad.Right(fmt.Sprintf("%.6f", r), pad)

	return fmt.Sprintf("%v | %v | %v |", paddedIter, paddedX, paddedR)
}

func shouldRun(xf XFlags, x float64) bool {
	if xf.delta > 0 {
		return x <= xf.end
	}

	return x >= xf.end
}

func Out(xf XFlags) {
	maxFlag := math.Max(xf.start, xf.end)
	maxColLen := len(fmt.Sprintf("%v", maxFlag)) + 4
	tableHead := fmt.Sprintf(
		"%s | %s | %s |",
		strpad.Right("#", maxColLen),
		strpad.Right("x", maxColLen),
		strpad.Right("result", maxColLen),
	)
	fmt.Println(tableHead)
	fmt.Println(strings.Repeat("—", len(tableHead)))

	i := 1

	for x := xf.start; shouldRun(xf, x); x += xf.delta {
		fmt.Println(row(i, x, maxColLen))
		i += 1
	}
}
