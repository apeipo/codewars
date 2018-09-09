package is_my_friend_cheating_test

import (
	. "github.com/apeipo/codewars/20180909_is-my-friend-cheating"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n uint64, exp [][2]uint64) {
	var ans = RemovNbV6(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("20180909IsMyFriendCheating", func() {
	It("should handle basic cases howMuch", func() {
		dotest(26, [][2]uint64{{15, 21}, {21, 15}})

		dotest(101, [][2]uint64{{55, 91}, {91, 55}})
		dotest(102, [][2]uint64{{70, 73}, {73, 70}})
		dotest(110, [][2]uint64{{70, 85}, {85, 70}})
	})

	It("should handle nil case", func() {
		dotest(100, nil)
	})
})
