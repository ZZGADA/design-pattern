package strategy

// Strategy 策略模式
type Strategy interface {
	run()
	initStrategy()
}

const (
	labelTreeKey = "label:tree"
)
