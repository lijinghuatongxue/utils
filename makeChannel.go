package utils

import (
	"sync"
)

func MakeChannel(numSum int, Func func()) {
	wg := sync.WaitGroup{}
	wg.Add(numSum)
	for i := 0; i < numSum; i++ {
		go func(i int) {
			// u's Func is here
			Func()
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// make channel to execute u's func
//func main() {
//	MakeChannel(100)
//}
