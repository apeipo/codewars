package n_linear

// https://www.codewars.com/kata/n-linear/cpp
// like twice linear

func Nlinear(m []int, n int) int {
	container := []int{1}
	mindex := map[int]int{}
	for _, v := range m {
		mindex[v] = 0
	}

	mval := map[int]int{}
	for len(container) < n+1 {
		for k, v := range mindex {
			mval[k] = k*container[v] + 1
		}
		// get min from mval
		minks, minv := getMin(mval)
		for _, k := range minks {
			mindex[k]++
		}
		container = append(container, minv)
	}
	return container[n]
}

func getMin(m map[int]int) ([]int, int) {
	mv := 1<<31 - 1
	mks := []int{}
	for k, v := range m {
		if v == mv {
			mks = append(mks, k)
		}
		// new min, need clean mks
		if v < mv {
			mks = append(mks[:0], k)
			mv = v
		}
	}
	return mks, mv
}

// c++ solutions
/*
#include <set>
#include <vector>

uint32_t n_linear(const std::set<uint32_t> &m_set, size_t n) {
    std::vector<uint32_t> container = {1};
    std::vector<uint32_t> m(m_set.begin(), m_set.end()), mv(m_set.size()), mi(m_set.size());
    for (int i = 0; i < m.size(); i++) {
        mv[i] = m[i] + 1;
        mi[i] = 0;
    }

    while (container.size() < n + 1) {
        const uint32_t minv = *std::min_element(mv.cbegin(), mv.cend());
        container.push_back(minv);
        for (int i = 0; i < m.size(); i++) {
            if (mv[i] == minv) {
                mi[i]++;
                mv[i] = m[i] * container[mi[i]] + 1;
            }
        }
    }
    return container[n];
}
*/
