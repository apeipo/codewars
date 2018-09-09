package consecutive_fib_number

//desc：https://www.codewars.com/kata/product-of-consecutive-fib-numbers/train/go

func fibSearch(n uint64, lastn uint64, search uint64) [3]uint64 {
	multN := n * lastn

	if multN == search {
		return [3]uint64{lastn, n, 1}
	}

	if multN > search {
		return [3]uint64{lastn, n, 0}
	}

	return fibSearch(n+lastn, n, search)
}

func fibSearchN(search uint64) [3]uint64 {
	var n, nextN, multN uint64 = 0, 1, 0
	for {
		multN = n * nextN
		if multN == search {
			return [3]uint64{n, nextN, 1}
		}

		if multN > search {
			return [3]uint64{n, nextN, 0}
		}

		n, nextN = nextN, n+nextN
	}
}

func ProductFib(prod uint64) [3]uint64 {
	return  fibSearchN(prod)
}


