package is_my_friend_cheating

import (
	"testing"
	"fmt"
)

func TestRemovNbV1(t *testing.T) {
	fmt.Printf("%+v\n", RemovNbV1(26)) // 15, 21;   21, 15
}

func TestRemovNbV2(t *testing.T) {
	fmt.Printf("%+v\n", RemovNbV2(26)) // 15, 21;   21, 15
}

func TestRemovNbV3(t *testing.T) {
	fmt.Printf("%+v\n", RemovNbV3(26)) // 15, 21;   21, 15
}

func TestRemovNbV4(t *testing.T) {
	fmt.Printf("%+v\n", RemovNbV4(26)) // 15, 21;   21, 15
}

func TestRemovNbV5(t *testing.T) {
	fmt.Printf("%+v\n", RemovNbV5(26)) // 15, 21;   21, 15
	fmt.Printf("%+v\n", RemovNbV5(10000)) // 15, 21;   21, 15
}

func TestRemovNbV6(t *testing.T) {
	fmt.Printf("%+v\n", RemovNbV6(26)) // 15, 21;   21, 15
	// fmt.Printf("%+v\n", RemovNbV5(10000)) // 15, 21;   21, 15
}


func BenchmarkRemovNbV1(b *testing.B) {
	b.ResetTimer()
	var n uint64 = 100000
	for i := 0; i < b.N; i++ {
		RemovNbV1(n)
	}
}

func BenchmarkRemovNbV2(b *testing.B) {
	b.ResetTimer()
	var n uint64 = 100000
	for i := 0; i < b.N; i++ {
		RemovNbV2(n)
	}
}

func BenchmarkRemovNbV3(b *testing.B) {
	b.ResetTimer()
	var n uint64 = 100000
	for i := 0; i < b.N; i++ {
		RemovNbV3(n)
	}
}

func BenchmarkRemovNbV4(b *testing.B) {
	b.ResetTimer()
	var n uint64 = 100000
	for i := 0; i < b.N; i++ {
		RemovNbV4(n)
	}
}

func BenchmarkRemovNbV5(b *testing.B) {
	b.ResetTimer()
	var n uint64 = 100000
	for i := 0; i < b.N; i++ {
		RemovNbV5(n)
	}
}

func BenchmarkRemovNbV6(b *testing.B) {
	b.ResetTimer()
	var n uint64 = 100000
	for i := 0; i < b.N; i++ {
		RemovNbV6(n)
	}
}

