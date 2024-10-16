package main

import "proxy"

func main() {
	// 创建代理对象
	_Proxy := proxy.NewProxy("real")

	// 代理对象调用主题接口的方法
	_Proxy.Say()
	_Proxy.Run()
}
