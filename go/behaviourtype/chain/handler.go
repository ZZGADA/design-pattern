package chain

import "fmt"

type Handler interface {
	HandleRequest(context *HandlerContext, msg string)
}

type ConcreteHandlerA struct{}

func (handler *ConcreteHandlerA) HandleRequest(context *HandlerContext, msg string) {
	//fmt.Println("Handle Request by handler A: ", msg+" A")
	if context.nextHandlerContext != nil {
		context.nextHandlerContext.Processed(msg + " A")
	}
}

type ConcreteHandlerB struct{}

func (handler *ConcreteHandlerB) HandleRequest(context *HandlerContext, msg string) {
	fmt.Println("Handle Request by handler B: ", msg+" B")
	if context.nextHandlerContext != nil {
		context.nextHandlerContext.Processed(msg + " end")
	}
}
