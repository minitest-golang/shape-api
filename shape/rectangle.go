package shape

import (
	"minitest/utils"
	"strconv"
)

type Rectangel struct {
	width  float64
	height float64
}

func CreateRectangle(edges []string) (*Rectangel, error) {
	if len(edges) != 2 {
		return nil, utils.ErrBadEdge
	}
	w, err := strconv.ParseFloat(edges[0], 64)
	if err != nil {
		return nil, utils.ErrBadEdge
	}
	h, err := strconv.ParseFloat(edges[1], 64)
	if err != nil {
		return nil, utils.ErrBadEdge
	}
	if w <= 0.0 || h <= 0.0 {
		return nil, utils.ErrBadEdge
	}
	return &Rectangel{
		width:  w,
		height: h,
	}, nil
}

func (r *Rectangel) Area() string {
	return strconv.FormatFloat(r.width*r.height, 'f', 4, 64)
}

func (r *Rectangel) Perimeter() string {
	return strconv.FormatFloat(2*(r.width+r.height), 'f', 4, 64)
}
