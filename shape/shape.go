package shape

import "minitest/utils"

const (
	TRIANGLE_SHAPE  = "triangle"
	RECTANGLE_SHAPE = "rectangle"
	SQUARE_SHAPE    = "square"
	DIAMOND_SHAPE   = "diamond"
)

type Shape interface {
	Area() string
	Perimeter() string
}

func CreateShape(shape string, edges []string) (Shape, error) {
	switch shape {
	case TRIANGLE_SHAPE:
		return CreateTriangle(edges)
	case RECTANGLE_SHAPE:
		return CreateRectangle(edges)
	case SQUARE_SHAPE:
		return CreateSquare(edges)
	case DIAMOND_SHAPE:
		return CreateDiamond(edges)
	}

	return nil, utils.ErrBadShape
}
