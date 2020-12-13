package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(data string) []int {
	strs := strings.Split(data, "\n")
	jolts := make([]int, len(strs))

	for i, s := range strs {
		jolt, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("error: canot convert to int: %s\n", s)
			os.Exit(1)
		}
		jolts[i] = jolt
	}

	return jolts
}

func main() {

	jolts := parseInput(data)
	sort.Ints(jolts)

	myJolt := jolts[len(jolts)-1] + 3
	jolts = append(jolts, myJolt)

	freq := make([]int, 3, 3)
	prev := 0
	for i := 0; i < len(jolts); i++ {
		if prev+1 == jolts[i] {
			freq[0]++
		} else if prev+2 == jolts[i] {
			freq[1]++
		} else if prev+3 == jolts[i] {
			freq[2]++
		} else {
			fmt.Printf("no adapter to connect after %d jolts. Closest is %d jolts\n", prev, jolts[i])
			os.Exit(1)
		}
		prev = jolts[i]
	}

	fmt.Println(freq)

	fmt.Println("freq(1) * freq(3) = ", freq[0]*freq[2])

}

var test1 = `16
10
15
5
1
11
7
19
6
12
4`

var test2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

var data = `153
17
45
57
16
147
39
121
75
70
85
134
128
115
51
139
44
65
119
168
122
72
105
31
103
89
154
114
55
25
48
38
132
157
84
71
113
143
83
64
109
129
120
100
151
79
125
22
161
167
19
26
118
142
4
158
11
35
56
18
40
7
150
99
54
152
60
27
164
78
47
82
63
46
91
32
135
3
108
10
159
127
69
110
126
133
28
15
104
138
160
98
90
144
1
2
92
41
86
66
95
12`
