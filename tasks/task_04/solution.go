package main

import (
	"math"
)

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	out := Stats{}

	// Проверка валидности данных
	if len(nums) <= 1 {
		return Stats{}
	}
	
	out.Count = len(nums) - 1
	out.Min = math.MaxInt
	out.Max = math.MinInt

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		out.Sum += diff
		if out.Max < diff {out.Max=diff}
		if out.Min > diff {out.Min=diff}
	}
	
	return out
}
