package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bradfitz/slice"
)

type N struct {
	num   int
	delay int
}

func main() {

	strs := strings.Split(data, "\n")

	nums := ParseNums(strs[1])

	fmt.Println(nums)

	fmt.Println("time = ", FirstTimestamp(nums))
}

func ParseNums(strs string) []N {
	strnums := strings.Split(strs, ",")
	nums := make([]N, 0)

	for i, sn := range strnums {
		n, err := strconv.Atoi(sn)
		if err != nil {
			continue
		}
		nums = append(nums, N{n, i})
	}

	slice.Sort(nums[:], func(i, j int) bool {
		return nums[i].num > nums[j].num
	})
	return nums
}

func FirstTimestamp(nums []N) int {
	time := 1
	dt := 1
	for {
		for i, n := range nums {
			if (time+n.delay)%n.num == 0 {
				dt *= n.num
				nums = append(nums[:i], nums[i+1:]...)
				break
			}
		}

		if len(nums) != 0 {
			time += dt
		} else {
			break
		}
	}
	return time
}

var test = `939
7,13,x,x,59,x,31,19`

var data = `1002394
13,x,x,41,x,x,x,37,x,x,x,x,x,419,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,23,x,x,x,x,x,29,x,421,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17`
