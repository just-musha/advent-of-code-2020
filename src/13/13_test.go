package main

import "testing"

func TestFirstTimestamp(t *testing.T) {
	testdata := []struct {
		name string
		in   string
		out  int
	}{
		{name: "test0", in: "7,13,x,x,59,x,31,19", out: 1068781},
		{name: "test1", in: "17,x,13,19", out: 3417},
		{name: "test2", in: "67,7,59,61", out: 754018},
		{name: "test3", in: "67,x,7,59,61", out: 779210},
		{name: "test4", in: "67,7,x,59,61", out: 1261476},
		{name: "test5", in: "1789,37,47,1889", out: 1202161486},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			nums := ParseNums(tt.in)

			res := FirstTimestamp(nums)

			if res != tt.out {
				t.Errorf("got %v, want %v", res, tt.out)
			}
		})
	}
}
