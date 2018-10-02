package n_linear_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/apeipo/codewars/20181001_n_linear"
)

func dotest(m []int, n int, exp int) {
	var ans = Nlinear(m, n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("20181001NLinear", func() {
	It("", func() {
		dotest([]int{2, 3}, 10, 22)
		dotest([]int{3, 2}, 10, 22)

		dotest([]int{5, 7, 8}, 10, 64)
		dotest([]int{5, 7, 8}, 11, 65)
	})
})
