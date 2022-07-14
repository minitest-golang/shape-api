package shape

import (
	"math"
	"minitest/utils"
	"strconv"
)

type Triangle struct {
	a float64
	b float64
	c float64
}

func CreateTriangle(edges []string) (*Triangle, error) {
	if len(edges) != 3 {
		return nil, utils.ErrBadEdge
	}
	a, err := strconv.ParseFloat(edges[0], 64)
	if err != nil {
		return nil, utils.ErrBadEdge
	}
	b, err := strconv.ParseFloat(edges[1], 64)
	if err != nil {
		return nil, utils.ErrBadEdge
	}
	c, err := strconv.ParseFloat(edges[2], 64)
	if err != nil {
		return nil, utils.ErrBadEdge
	}
	if a <= 0.0 || b <= 0.0 || c <= 0.0 {
		return nil, utils.ErrBadEdge
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return nil, utils.ErrBadEdge
	}
	return &Triangle{
		a: a,
		b: b,
		c: c,
	}, nil
}

func (t *Triangle) Area() string {
	S := (t.a + t.b + t.c) / 2
	T := S * (S - t.a) * (S - t.b) * (S - t.c)
	A := math.Sqrt(T)
	return strconv.FormatFloat(A, 'f', 4, 64)
}

func (t *Triangle) Perimeter() string {
	return strconv.FormatFloat(t.a+t.b+t.c, 'f', 4, 64)
}
