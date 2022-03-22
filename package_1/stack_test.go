package main

import (
	"github.com/multipackage_1/package_3"
	"reflect"
	"testing"
)

func Test_makeStack(t *testing.T) {
	tests := []struct {
		name string
		want *package_3.Stack
	}{
		{"1", package_3.InitializeStack()},
		{"2", package_3.InitializeStack()},
		{"3", package_3.InitializeStack()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
