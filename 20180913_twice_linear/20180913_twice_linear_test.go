package twice_linear_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/apeipo/codewars/20180913_twice_linear"
)

func dotest(n int, exp int) {
	var ans = DblLinearV2(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("20180913TwiceLinear", func() {
		 // dotest(10, 22)
		dotest(50, 175)
		// dotest(100, 447)
		// dotest(500, 3355)
		dotest(1000, 8488)
	})
})