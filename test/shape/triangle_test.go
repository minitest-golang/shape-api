package shape_test

import (
	"minitest/shape"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCreateTriangle(t *testing.T) {
	convey.Convey("Test@CreateTriangle - Create Triangle OK", t, func() {
		s, err := shape.CreateTriangle([]string{
			"3",
			"4.0",
			"5.00",
		})
		convey.So(s, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)
	})
	convey.Convey("Test@CreateTriangle - Create Triangle failed due to invalid edges length", t, func() {
		s, err := shape.CreateTriangle([]string{
			"11",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateTriangle - Create Triangle failed due to invalid edges length", t, func() {
		s, err := shape.CreateTriangle([]string{
			"11",
			"22",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateTriangle - Create Triangle failed due to invalid edges length", t, func() {
		s, err := shape.CreateTriangle([]string{
			"11",
			"22",
			"33.0",
			"44.1",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateTriangle - Create Triangle failed due to bad edges value", t, func() {
		s, err := shape.CreateTriangle([]string{
			"a3",
			"4.0",
			"5.00",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateTriangle - Create Triangle failed due to bad edges value", t, func() {
		s, err := shape.CreateTriangle([]string{
			"3",
			"0.0",
			"5.00",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateTriangle - Create Triangle failed due to negative edges value", t, func() {
		s, err := shape.CreateTriangle([]string{
			"3",
			"4.0",
			"-5.00",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})

	convey.Convey("Test@CreateTriangle - Create Triangle failed due to bad edges value", t, func() {
		s, err := shape.CreateTriangle([]string{
			"3",
			"4.0",
			"9.00",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func TestTriangleArea(t *testing.T) {
	convey.Convey("Test@TriangleArea - Triangle Area OK", t, func() {
		s, _ := shape.CreateTriangle([]string{
			"3.0",
			"4.00",
			"5.000",
		})
		a := s.Area()
		convey.So(a, convey.ShouldEqual, "6.0000")
	})
}

func TestTrianglePerimeter(t *testing.T) {
	convey.Convey("Test@TrianglePerimeter - Triangle Perimeter OK", t, func() {
		s, _ := shape.CreateTriangle([]string{
			"3.0",
			"4.00",
			"5.000",
		})
		a := s.Perimeter()
		convey.So(a, convey.ShouldEqual, "12.0000")
	})
}
