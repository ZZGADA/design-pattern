package main

import (
	"chain"
	"strconv"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	handlerA := &chain.ConcreteHandlerA{}
	handlerB := &chain.ConcreteHandlerB{}

	wg.Add(10)
	for i := 0; i <= 10; i++ {
		go func() {
			defer wg.Done()
			contextB := chain.NewHandlerContext(handlerB, nil)
			contextA := chain.NewHandlerContext(handlerA, &contextB)

			contextA.Processed("hello world " + strconv.Itoa(i))
		}()
	}
	wg.Wait()

}
