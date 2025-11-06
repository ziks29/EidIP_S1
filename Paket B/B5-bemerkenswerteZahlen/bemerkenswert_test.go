package main

import "testing"

func TestBemerkenswert(t *testing.T) {

	tests := [...]struct {
		start int
		want  int
	}{
		{0, 1},
		{1, 6},
		{6, 35},
		{35, 204},
		{204, 1189},
		{1189, 6930},
		{6930, 40391},
		// {40391, 235416}, dauert zu lange
		//{235416, 1372105},
		//{1372105, 7997214},
		//{7997214, 46611179},
		//{46611179, 271669860},
		//{271669860, 1583407981},
		//{1583407981, 9228778026},
		//{9228778026, 53789260175},
		//{53789260175, 313506783024},
		//{313506783024, 1827251437969},
		//{1827251437969, 10650001844790},
	}
	for _, tt := range tests {
		got := bemerkenswert(tt.start)
		if got != tt.want {
			t.Errorf("bemerkenswert(%v) = %v, want %v", tt.start, got, tt.want)
		}
	}
}
