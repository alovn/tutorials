package main

import (
	"testing"
)

func Test_create(t *testing.T) {
	fs := create()
	for _, f := range fs {
		f()
	}
}
