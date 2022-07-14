package shape

import (
	"minitest/utils"
	"strconv"
)

type Square struct {
	side float64
}

func CreateSquare(edges []string) (*Square, error) {
	if len(edges) != 1 {
		return nil, utils.ErrBadEdge
	}
	s, err := strconv.ParseFloat(edges[0], 64)
	if err != nil {
		return nil, utils.ErrBadEdge
	}
	if s <= 0.0 {
		return nil, utils.ErrBadEdge
	}
	return &Square{
		side: s,
	}, nil
}

func (s *Square) Area() string {
	return strconv.FormatFloat(s.side*s.side, 'f', 4, 64)
}

func (s *Square) Perimeter() string {
	return strconv.FormatFloat(4*s.side, 'f', 4, 64)
}
