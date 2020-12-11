package main

import (
	"strings"
	"testing"
)

var data1 = `.............
.L.L.#.#.#.#.
.............`

var data2 = `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`

var data3 = `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`

func TestCalcBusySeats(t *testing.T) {
	var testdata = []struct {
		name   string
		in     string
		ti, tj int
		out    int
	}{
		{
			name: "test1",
			in:   data1,
			ti:   1,
			tj:   1,
			out:  0,
		},
		{
			name: "test2",
			in:   data2,
			ti:   3,
			tj:   3,
			out:  0,
		},
		{
			name: "test3",
			in:   data3,
			ti:   4,
			tj:   3,
			out:  8,
		},
		{
			name: "test4",
			in:   data1,
			ti:   0,
			tj:   5,
			out:  1,
		},
		{
			name: "test4",
			in:   data1,
			ti:   0,
			tj:   6,
			out:  2,
		},
		{
			name: "test5",
			in:   data1,
			ti:   0,
			tj:   0,
			out:  0,
		},
		{
			name: "test5",
			in:   data3,
			ti:   8,
			tj:   8,
			out:  2,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			strs := strings.Split(tt.in, "\n")

			xmax := len(strs)
			ymax := len(strs[0])

			places := make([][]int, xmax)
			for i := range places {
				places[i] = make([]int, ymax)
			}
			initSeats(places, xmax, ymax, strs)

			s := countBusyAround(places, tt.ti, tt.tj, xmax, ymax)
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}
