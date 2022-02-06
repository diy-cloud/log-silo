package signal

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const signalNum = 1

var sigs = make(chan os.Signal, 1)
var done = make(chan struct{}, signalNum)
var once = new(sync.Once)

func NewTerminate() <-chan struct{} {
	once.Do(func() {
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			sig := <-sigs
			fmt.Println()
			fmt.Println(sig)
			for i := 0; i < signalNum; i++ {
				done <- struct{}{}
			}
		}()
	})

	return done
}
