package shape_test

import (
	"minitest/shape"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCreateDiamond(t *testing.T) {
	convey.Convey("Test@CreateDiamond - Create Diamond OK", t, func() {
		s, err := shape.CreateDiamond([]string{
			"11",
			"22.2",
		})
		convey.So(s, convey.ShouldNotBeNil)
		convey.So(err, convey.ShouldBeNil)
	})
	convey.Convey("Test@CreateDiamond - Create Diamond failed due to invalid diagonal length", t, func() {
		s, err := shape.CreateDiamond([]string{
			"11",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateDiamond - Create Diamond failed due to invalid diagonal length", t, func() {
		s, err := shape.CreateDiamond([]string{
			"11",
			"22",
			"33",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateDiamond - Create Diamond failed due to bad diagonal value", t, func() {
		s, err := shape.CreateDiamond([]string{
			"11a",
			"22.2",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateDiamond - Create Diamond failed due to bad diagonal value", t, func() {
		s, err := shape.CreateDiamond([]string{
			"11",
			"0",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
	convey.Convey("Test@CreateDiamond - Create Diamond failed due to negative diagonal value", t, func() {
		s, err := shape.CreateDiamond([]string{
			"11",
			"-1.00",
		})
		convey.So(s, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func TestDiamondArea(t *testing.T) {
	convey.Convey("Test@DiamondArea - Diamond Area OK", t, func() {
		s, _ := shape.CreateDiamond([]string{
			"2",
			"4",
		})
		a := s.Area()
		convey.So(a, convey.ShouldEqual, "4.0000")
	})
}

func TestDiamondPerimeter(t *testing.T) {
	convey.Convey("Test@DiamondPerimeter - Diamond Perimeter OK", t, func() {
		s, _ := shape.CreateDiamond([]string{
			"6",
			"8",
		})
		a := s.Perimeter()
		convey.So(a, convey.ShouldEqual, "20.0000")
	})
}
