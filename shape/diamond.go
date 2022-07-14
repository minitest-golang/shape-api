package shape

import (
	"math"
	"minitest/utils"
	"strconv"
)

type Diamond struct {
	diagonal_1 float64
	diagonal_2 float64
}

func CreateDiamond(edges []string) (*Diamond, error) {
	if len(edges) != 2 {
		return nil, utils.ErrBadEdge
	}
	d1, err := strconv.ParseFloat(edges[0], 64)
	if err != nil {
		return nil, utils.ErrBadEdge
	}
	d2, err := strconv.ParseFloat(edges[1], 64)
	if err != nil {
		return nil, utils.ErrBadEdge
	}
	if d1 <= 0.0 || d2 <= 0.0 {
		return nil, utils.ErrBadEdge
	}
	return &Diamond{
		diagonal_1: d1,
		diagonal_2: d2,
	}, nil
}

func (d *Diamond) Area() string {
	return strconv.FormatFloat((d.diagonal_1*d.diagonal_2)/2, 'f', 4, 64)
}

func (d *Diamond) Perimeter() string {
	P := 2 * math.Sqrt((d.diagonal_1*d.diagonal_1)+(d.diagonal_2*d.diagonal_2))
	return strconv.FormatFloat(P, 'f', 4, 64)
}
