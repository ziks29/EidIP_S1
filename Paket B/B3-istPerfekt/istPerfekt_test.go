package main

import (
	"testing"
)

func TestIstPerfekt(t *testing.T) {
	tests := [...]struct {
		z    int
		want bool
	}{
		{6, true},
		{28, true},
		{496, true},
		{8128, true},
		{33550336, true},
		//{8589869056, true}, dauert zu lange...
		//{137438691328, true},         "
		//{2305843008139952128, true},  "
		{10, false},
		{11, false},
		{12, false},
		{13, false},
		{14, false},
		{15, false},
		{16, false},
		{17, false},
		{18, false},
		{19, false},
		{20, false},
		{21, false},
		{22, false},
		{33, false},
		{42, false},
		{66, false},
		{88, false},
		{99, false},
		{100, false},
		{112, false},
		{152, false},
		{172, false},
		{181, false},
		{470, false},
		{200, false},
		{213, false},
		{224, false},
		{299, false},
		{351, false},
		{372, false},
		{531, false},
		{531, false},
		{513, false},
		{987, false},
		{1810, false},
		{1971, false},
		{2002, false},
		{2138, false},
		{2247, false},
		{1633, false},
		{8200, false},
		{9472, false},
		{10000, false},
		{11111, false},
		{111112, false},
		{1111123, false},
		{33550337, false},
		{33550335, false},
		{33553377, false},
	}
	for _, tt := range tests {
		got := istPerfekt(tt.z)
		if got != tt.want {
			t.Errorf("istPerfekt(%v) = %v, want %v", tt.z, got, tt.want)
		}
	}
}
