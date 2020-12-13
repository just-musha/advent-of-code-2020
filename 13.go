package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	strs := strings.Split(data, "\n")

	start, err := strconv.Atoi(strs[0])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	strnums := strings.Split(strs[1], ",")
	nums := make([]int, 0)
	for _, sn := range strnums {
		n, err := strconv.Atoi(sn)
		if err != nil {
			continue
		}
		nums = append(nums, n)
	}

	minTime := int(^uint(0) >> 1)
	busID := -1

	for _, n := range nums {
		mult := (start + n) / n
		busStart := mult * n

		if busStart < minTime {
			minTime = busStart
			busID = n
		}
	}
	fmt.Println("Earliest bus start = ", minTime)
	fmt.Println("Time to wait = ", minTime-start)
	fmt.Println("Answer = ", (minTime-start)*busID)
}

var test = `939
7,13,x,x,59,x,31,19`

var data = `1002394
13,x,x,41,x,x,x,37,x,x,x,x,x,419,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,23,x,x,x,x,x,29,x,421,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17`
