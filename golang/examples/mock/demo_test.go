package demo

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetString(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockHTTP := NewMockHTTP(mockCtl)
	mockHTTP.EXPECT().Get().Return("mock demo")
	getStr := GetString(mockHTTP)
	t.Log(getStr)
	if getStr != "mock demo" {
		t.Error("GetString ERROR")
	}
}
