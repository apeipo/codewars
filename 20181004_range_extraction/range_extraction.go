package range_extraction

import (
	"bytes"
	"fmt"
	"strings"
	"strconv"
)

func formatSub(buf *bytes.Buffer, list []int, s int, e int) {
	if e-s >= 2 {
		buf.WriteString(fmt.Sprintf("%d-%d,", list[s], list[e]))
		return
	}

	for i := s; i <= e && i < len(list); i++ {
		buf.WriteString(fmt.Sprintf("%d,", list[i]))
	}
}

// https://www.codewars.com/kata/range-extraction
func Solution(list []int) string {
	pre, cur := 0, 1
	buf := &bytes.Buffer{}
	for cur < len(list) {
		if list[cur]-list[cur-1] == 1 {
			cur++
			continue
		}
		if list[cur]-list[cur-1] > 1 {
			// fmt.Println(cur, pre)
			formatSub(buf, list, pre, cur-1)
			pre = cur
			cur++
		}
	}
	formatSub(buf, list, pre, cur-1)
	return strings.TrimRight(buf.String(), ",")
}

func SolutionV2(list []int) string {
	l := len(list)
	s := ""
	for i := 0; i < l; {
		s += strconv.Itoa(list[i])
		j := i+1;
		for j < l && list[j]-list[j-1] == 1 {
			j++
		}
		if j - i > 2 {
			s += "-" + strconv.Itoa(list[j-1])
			i = j
		} else {
			i++
		}
		s += ","
	}
	return strings.TrimRight(s, ",")
}
