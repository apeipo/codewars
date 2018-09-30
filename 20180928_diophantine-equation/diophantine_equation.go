package diophantine_equation

import (
	"math"
)

// https://www.codewars.com/kata/diophantine-equation/train/go
//  find all integers x, y (x >= 0, y >= 0) solutions of a diophantine equation :
// x^2 - 4 * y^2 = n

func SolequaV1(n int) [][]int {
	// your code
	r := [][]int{}
	for x := n - 1; x >= 0; x-- {
		y := int(float64(x-n) / 2)
		if y < 0 {
			y = 0
		}
		for ; 2*y < x; y++ {
			if x*x-4*y*y == n {
				r = append(r, []int{x, y})
			}
		}
	}
	return r
}

// reduce the caculate
func SolequaV2(n int) [][]int {
	// your code
	r := [][]int{}
	for x := n - 1; x >= 0; x-- {
		t := x*x - n
		if t%4 != 0 {
			continue
		}
		y := int(math.Sqrt(float64(t) / 4))
		if x*x-y*y*4 == n {
			r = append(r, []int{x, y})
		}
	}
	return r
}

// reduce loop
func SolequaV3(n int) [][]int {
	// your code
	r := [][]int{}
	low := int(math.Sqrt(float64(n)))
	for x := int(n/2) + 1; x >= low; x-- {
		t := x*x - n
		if t < 0 || t%4 != 0 {
			continue
		}
		y := int(math.Sqrt(float64(t) / 4))
		if 2*y > x {
			continue
		}
		if (x-2*y)*(x+2*y) == n {
			r = append(r, []int{x, y})
		}
	}
	return r
}

// from (x-2y)*(x+2y) == n
// we let a = (x-2y), b=(x+2y)
// then get:
// a*b = n
// x = (a+b)/2
// y = (b-a)/4
// 1. x>=0 && y>=0  =>  "b>=a"  => "a < sqrt(n)"
// 2. x, y is positive number => "(a+b)%2 == 0 and (b-a)%4==0"
func SolequaV4(n int) [][]int {
	// your code
	r := [][]int{}
	//i*j == n
	low := int(math.Sqrt(float64(n)))
	for i := 1; i <= low; i++ {
		if n%i != 0 {
			continue
		}
		j := n / i
		if j < i {
			continue
		}
		if ((i+j)%2 == 0) && ((j-i)%4 == 0) {
			r = append(r, []int{(i+j)/2, (j-i)/4})
		}
	}
	return r
}
