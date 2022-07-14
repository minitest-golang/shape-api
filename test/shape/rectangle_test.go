package shape_test

import (
	"minitest/shape"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCreateRectangle(t *testing.T) {
	convey.Convey("Test@CreateRectangle - Create Rectangle OK", t, func() {
		s, err := shape.CreateRectangle([]string{
			"11",
			"22.2",
		})
		convey.So(s, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)
	})
	convey.Convey("Test@CreateRectangle - Create Rectangle failed due to invalid edges length", t, func() {
		s, err := shape.CreateRectangle([]string{
			"11",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateRectangle - Create Rectangle failed due to invalid edges length", t, func() {
		s, err := shape.CreateRectangle([]string{
			"11",
			"22",
			"33",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateRectangle - Create Rectangle failed due to bad edges value", t, func() {
		s, err := shape.CreateRectangle([]string{
			"11a",
			"22.2",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateRectangle - Create Rectangle failed due to bad edges value", t, func() {
		s, err := shape.CreateRectangle([]string{
			"11",
			"0",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateRectangle - Create Rectangle failed due to negative edges value", t, func() {
		s, err := shape.CreateRectangle([]string{
			"11",
			"-1.0",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func TestRectangleArea(t *testing.T) {
	convey.Convey("Test@RectangleArea - Rectangle Area OK", t, func() {
		s, _ := shape.CreateRectangle([]string{
			"2",
			"4",
		})
		a := s.Area()
		convey.So(a, convey.ShouldEqual, "8.0000")
	})
}

func TestRectanglePerimeter(t *testing.T) {
	convey.Convey("Test@RectanglePerimeter - Rectangle Perimeter OK", t, func() {
		s, _ := shape.CreateRectangle([]string{
			"2",
			"4",
		})
		a := s.Perimeter()
		convey.So(a, convey.ShouldEqual, "12.0000")
	})
}
