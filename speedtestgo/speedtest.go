package speedtestgo

import (
	"fmt"
	"time"
)

func SpeedTest() {

	n := 1000000
	test := []int{}
	testSlice := make([]int, 0, n)

	fmt.Printf("Total time taken  without preallocation %v\n", timeLoop(test, n))
	fmt.Printf("Total time taken  with preallocation %v\n", timeLoop(testSlice, n))

}

func timeLoop(slice[]int, n int)time.Duration{
	t0:=time.Now()
	for len(slice)<n{
		slice = append(slice, 1)
	}
	return time.Since(t0)
}