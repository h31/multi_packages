package package_3

import (
	"reflect"
	"testing"
)

func Test_makeStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		{"1", InitializeStack()},
		{"2", InitializeStack()},
		{"3", InitializeStack()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
