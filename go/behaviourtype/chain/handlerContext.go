package chain

type HandlerContext struct {
	currentHandler     Handler
	nextHandlerContext *HandlerContext
}

func NewHandlerContext(handler Handler, nextHandlerContext *HandlerContext) HandlerContext {
	return HandlerContext{
		currentHandler:     handler,
		nextHandlerContext: nextHandlerContext,
	}
}

// Processed 由上下文控制handler的执行
func (hc *HandlerContext) Processed(msg string) {
	if hc.currentHandler != nil {
		hc.currentHandler.HandleRequest(hc, msg)
	}
}
