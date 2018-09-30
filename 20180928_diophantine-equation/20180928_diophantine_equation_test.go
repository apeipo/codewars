package diophantine_equation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/apeipo/codewars/20180928_diophantine-equation"
)

func dotest(n int, exp [][]int) {
	var ans =diophantine_equation.SolequaV4(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("20180928DiophantineEquation", func() {
	It("should handle basic cases", func() {
		dotest(5, [][]int{{3, 1}})
		dotest(12, [][]int{{4, 1}})
		dotest(13, [][]int{{7, 3}})
		dotest(16, [][]int{{4, 0}})
		dotest(9005, [][]int{{4503, 2251}, {903, 449}})
		dotest(9008, [][]int{{1128, 562}})
		dotest(90005, [][]int{{45003, 22501}, {9003, 4499}, {981, 467}, {309, 37}})
		dotest(90002, [][]int{})
		// dotest(9000000041, [][]int{})
	})
})
