package range_extraction_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/apeipo/codewars/20181004_range_extraction"
	. "github.com/onsi/gomega"
)

func dotest(arr []int, expect string) {
	s := SolutionV2(arr)
	Expect(s).To(Equal(expect))
}

var _ = Describe("Range Extraction", func() {
	It("should match the example", func() {
		dotest([]int{-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20}, "-6,-3-1,3-5,7-11,14,15,17-20")
		dotest([]int{}, "")

	})
})