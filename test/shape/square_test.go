package shape_test

import (
	"minitest/shape"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCreateSquare(t *testing.T) {
	convey.Convey("Test@CreateSquare - Create Square OK", t, func() {
		s, err := shape.CreateSquare([]string{
			"11",
		})
		convey.So(s, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)
	})
	convey.Convey("Test@CreateSquare - Create Square failed due to invalid edges length", t, func() {
		s, err := shape.CreateSquare([]string{})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateSquare - Create Square failed due to invalid edges length", t, func() {
		s, err := shape.CreateSquare([]string{
			"11",
			"22",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateSquare - Create Square failed due to bad edges value", t, func() {
		s, err := shape.CreateSquare([]string{
			"11a",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateSquare - Create Square failed due to bad edges value", t, func() {
		s, err := shape.CreateSquare([]string{
			"0",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateSquare - Create Square failed due to negative edges value", t, func() {
		s, err := shape.CreateSquare([]string{
			"-1.00",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func TestSquareArea(t *testing.T) {
	convey.Convey("Test@SquareArea - Square Area OK", t, func() {
		s, _ := shape.CreateSquare([]string{
			"2",
		})
		a := s.Area()
		convey.So(a, convey.ShouldEqual, "4.0000")
	})
}

func TestSquarePerimeter(t *testing.T) {
	convey.Convey("Test@SquarePerimeter - Square Perimeter OK", t, func() {
		s, _ := shape.CreateSquare([]string{
			"2",
		})
		a := s.Perimeter()
		convey.So(a, convey.ShouldEqual, "8.0000")
	})
}
