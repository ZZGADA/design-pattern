package main

import (
	"singleton"
	"sync"
)

func main() {
	s := singleton.Singleton{}
	wg := new(sync.WaitGroup)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			s.Get()
		}()
	}
	wg.Wait()

}
