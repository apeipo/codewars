package is_my_friend_cheating

import "fmt"

// desc https://www.codewars.com/kata/is-my-friend-cheating
// give a number n， find an number pair <a, b>,  a * b = sum(1, n) - a - b

func RemovNbV1(m uint64) [][2]uint64 {
	// your code
	var i, j, sum uint64
	res := make([][2]uint64, 0)
	sum = (1 + m) * m / 2
	for i = 1; i < m; i++ {
		for j = 1; j < m; j++ {
			if sum-i-j == i*j {
				res = append(res, [2]uint64{i, j})
			}
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

func RemovNbV2(m uint64) [][2]uint64 {
	// your code
	var i, j, sum uint64
	res := make([][2]uint64, 0)
	sum = (1 + m) * m / 2
	for i = 1; i < m; i++ {
		for j = i; j < m; j++ {
			if sum-i-j == i*j {
				res = append(res, [2]uint64{i, j}, [2]uint64{j, i})
			}
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

func RemovNbV3(m uint64) [][2]uint64 {
	// your code
	var i, j, sum uint64
	res := make([][2]uint64, 0, 10)
	sum = (1 + m) * m / 2
	for i = 1; i < m; i++ {
		if i*m+i+m <= sum {
			continue
		}
		for j = i; j < m; j++ {
			if sum-i-j == i*j {
				res = append(res, [2]uint64{i, j}, [2]uint64{j, i})
			}
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

func RemovNbV4(m uint64) [][2]uint64 {
	// your code
	var i, j, sum uint64
	res := make([][2]uint64, 0, m/2)
	sum = (1 + m) * m / 2
	for i = 1; i < m; i++ {
		for j = m; j >= i; j-- {
			if i*j+i+j < sum {
				break
			}
			if sum == i*j+i+j {
				res = append(res, [2]uint64{i, j}, [2]uint64{j, i})
			}
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

func RemovNbV5(m uint64) [][2]uint64 {
	// your code
	var i, j, sum uint64
	res := make([][2]uint64, 0, m/2)
	sum = (1 + m) * m / 2
	cal := 0
	for i = 1; i < m; i++ {
		for j = m; j >= i; j-- {
			cal++
			t := i*j + i + j
			if t < sum {
				break
			}
			if t == sum {
				res = append(res, [2]uint64{i, j}, [2]uint64{j, i})
			}
		}
	}
	fmt.Printf("loopcn:%d\n", cal)
	if len(res) == 0 {
		return nil
	}
	return res
}

func RemovNbV6(m uint64) [][2]uint64 {
	// your code
	var i, j, sum uint64
	var res [][2]uint64
	sum = (1 + m) * m / 2
	for i = 1; i < m; i++ {
		//inner loop is not needed
		j = (sum - i) / (i + 1)
		if (j <= m) && (i*j+i+j == sum ){
			res = append(res, [2]uint64{i, j})
		}
	}
	return res
}

// #include <vector>
//
// using namespace std;
//
// class RemovedNumbers
// {
// public:
// static vector<vector<long long>> removNb(long long n);
// };
//
// vector<vector<long long>> RemovedNumbers::removNb(long long n) {
// long long i, j , sum = 0;
// sum = (1+n)*n/2;
// vector<vector<long long>> r;
// for (i=1; i < n; i++) {
// for (j=n; j>=i; j--) {
// if (i*j+i+j < sum) {
// break;
// }
// if (i*j+i+j == sum) {
// r.push_back(vector<long long>{i, j});
// r.push_back(vector<long long>{j, i});
// }
// }
// }
// return r;
// }
