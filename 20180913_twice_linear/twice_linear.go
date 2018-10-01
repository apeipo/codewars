package twice_linear

import (
	"sort"
)

func DblLinearV1(n int) int {
	numDict := make(map[int]bool)
	nums := make([]int, 0)
	part := []int{1}
	nMax := 0 // max num in continer
	enough := false
	for {
		temp := make([]int, 0, 2*len(part))
		partMin := 1<<31 - 1
		for _, i := range part {
			if enough {
				if i < nMax {
					if _, ok := numDict[i]; !ok {
						numDict[i] = true
						nums = append(nums, i)
					}
				}
			} else {
				if _, ok := numDict[i]; !ok {
					numDict[i] = true
					nums = append(nums, i)
					if i > nMax {
						nMax = i
					}
					if len(numDict) == n {
						enough = true
					}
				}
			}

			if i < partMin {
				partMin = i
			}
			temp = append(temp, 2*i+1, 3*i+1)
		}

		if partMin > nMax {
			break
		}
		part = temp
	}
	sort.Ints(nums)
	return nums[n]
}

func DblLinearV2(n int) int {
	x, y := 0, 0
	container := make([]int, 0)
	container = append(container, 1)
	for len(container) < n + 1 {
		nexta := 2*container[x] + 1
		nextb := 3*container[y] + 1
		switch  {
		case nexta < nextb:
			container = append(container, nexta)
			x++
		case nexta > nextb:
			container = append(container, nextb)
			y++
		case nexta == nextb:
			container = append(container, nexta)
			x++
			y++
		}
	}
	return container[n]
}