package test

import (
	"blc/BLC"
	"reflect"
	"testing"
)

// 测试IntToHex函数
func TestIntToHex(t *testing.T) {
	tests := []struct {
		name string
		num  int64
		want []byte
	}{
		{"test1", 0, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{"test2", 1, []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{"test3", 256, []byte{0, 0, 0, 0, 0, 0, 1, 0}},
		{"test4", -1, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BLC.IntToHex(tt.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToHex(%d) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
