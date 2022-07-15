package main

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMockFunc(t *testing.T) {
	convey.Convey("TestMockFunc1", t, func() {
		var p1 = gomonkey.ApplyFunc(netWorkFunc, func(a, b int) (int, error) {
			fmt.Println("in mock function")
			return a + b, nil
		})
		defer p1.Reset()

		sum, err := logicFunc(10, 20)
		convey.So(sum, convey.ShouldEqual, 30)
		convey.So(err, convey.ShouldBeNil)
	})

}
