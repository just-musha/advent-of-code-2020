package main

import (
	"strings"
	"testing"
)

func TestDay7(t *testing.T) {
	var testdata = []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "test",
			in:   test,
			out:  126,
		},
		{
			name: "test1",
			in:   test1,
			out:  42,
		},
		{
			name: "test2",
			in:   test2,
			out:  30,
		},
		{
			name: "test3",
			in:   test3,
			out:  50,
		},
		{
			name: "test4",
			in:   test4,
			out:  77,
		},
		// {
		//	name: "data",
		// 	in:  data,
		// 	out: 1000,
		// },
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {

			bags := make(map[string]NextBags)

			strs := strings.Split(tt.in, "\n")
			for _, s := range strs {
				parseBagsInfo(bags, s)
			}

			s := countAllBags(bags, "shiny gold")
			if s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}
		})
	}
}
