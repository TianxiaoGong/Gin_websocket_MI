package utils

import (
	"time"
)

type TimerFunc func(interface{}) bool

// Timer delay首次延迟 tick间隔 fun定时执行的方法 param方法的参数
func Timer(delay, tick time.Duration, fun TimerFunc, param interface{}) {
	go func() {
		if fun == nil {
			return
		}
		t := time.NewTimer(delay)
		for {
			select {
			case <-t.C:
				if fun(param) == false {
					return
				}
				t.Reset(tick)
			}
		}
	}()
}
