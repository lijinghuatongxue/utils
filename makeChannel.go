package main

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

func Task() {
	logrus.Info(runtime.NumGoroutine())
}

func MakeChannel(numSum int) {
	wg := sync.WaitGroup{}
	wg.Add(numSum)
	for i := 0; i < numSum; i++ {
		go func(i int) {
			// u's Func is here
			Task()
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// make channel to execute u's func
//func main() {
//	MakeChannel(100)
//}
