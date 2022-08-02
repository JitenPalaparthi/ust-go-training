package demo_test

import (
	"demomock/demo"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_Reverse(t *testing.T) {
	ctl := gomock.NewController(t)
	//defer ctl.Finish()
	mockReverse := demo.NewMockReverser(ctl)

	expectedOutout := "olleH"
	input := "Hello"
	//gomock.Eq()
	mockReverse.EXPECT().Reverse(gomock.Eq(input)).Return(expectedOutout).MaxTimes(1)

}

func Test_Count(t *testing.T) {
	ctl := gomock.NewController(t)
	//defer ctl.Finish()
	mockReverse := demo.NewMockReverser(ctl)

	expectedOutout := 5
	input := "Hello"
	//gomock.Eq()
	mockReverse.EXPECT().Count(input).Return(expectedOutout).MaxTimes(1)

}
