package runner

import (
	"fmt"
	"time"
)

func RunInt(f func(int) int, arg int, reps int) {
	var result int
	start := time.Now()
	for i := 0; i < reps; i++ {
		result = f(arg)
	}
	end := time.Now()
	fmt.Printf("%v\nRunning %v times took %v.\n", result, reps, end.Sub(start))
}

func RunInt64(f func(int64) int64, arg int64, reps int) {
	var result int64
	start := time.Now()
	for i := 0; i < reps; i++ {
		result = f(arg)
	}
	end := time.Now()
	fmt.Printf("%v\nRunning %v times took %v.\n", result, reps, end.Sub(start))
}
