package main

import "adapter"

func main() {
	// 实例化一个adaptee对象 （现实世界理解成一个第三方的服务或者包）
	adapteeImpl := adapter.AdapteeImpl{}
	adaptor := adapter.NewAdapter(adapteeImpl)

	// 调用适配器适配后的方法
	adaptor.FitRequest()
}
