package main

import "strategy"

func main() {
	// 实例化待执行的策略
	strategyUsingCash := &strategy.CashPayStrategy{}
	strategyUsingCreditCart := &strategy.CreditPayStrategy{}

	// 实例化上下文（理解为执行器）
	contextExecutor := strategy.Context{}

	// 执行具体的策略
	contextExecutor.SetStrategy(strategyUsingCreditCart)
	contextExecutor.ExecuteStrategy(100)

	contextExecutor.SetStrategy(strategyUsingCash)
	contextExecutor.ExecuteStrategy(100)
}
