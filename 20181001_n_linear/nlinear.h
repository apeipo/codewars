//
// Created by 林仙龙 on 2018/10/2.
//

#ifndef CTEST_NLINEAR_V3_H
#define CTEST_NLINEAR_V3_H

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


#endif //CTEST_NLINEAR_V3_H
