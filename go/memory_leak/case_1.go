package main

func main() {
	// case_1
	{
		// 创建一个切片并传递给 test 函数
		// test([]int{1, 2, 3})
		// 此时切片 a 包含了 [1, 2, 3, 4, 5, 6]
		// b 数组将无法被垃圾回收，因为它与 a 共享底层数组的地址
		// 如果没有其他地方引用切片 a，那么 a 将会被垃圾回收
	}
}

var a []int

// a 为全局变量的切片 b和a公用内存地址 a不gc 那么b也无发gc
// 因此会发生内存泄漏
func test(b []int) {
	a = b[:3]
	// 此处对切片 a 进行追加操作
	a = append(a, 4, 5, 6)
	return
}

var s0 string // a package-level variable

// A demo purpose function.
// 如果 s1 的大小非常大（例如，几十 MB 或更多），且 s0 只使用其前 50 字节，整个大内存块会因为 s0 的引用而无法回收，
// 直到 s0 被重新赋值或程序退出。这种情况会导致内存浪费，可以视为一种内存泄漏。
/**
如果 f 只被调用一次，且 s1 的大小适中（例如，几 KB），内存浪费可能不显著，不会被认为是严重的内存泄漏。
如果 f 被频繁调用，且每次传入的 s1 是一个非常大的字符串（例如，读取大文件或网络数据），而 s0 只保留前 50 字节，内存浪费会更明显，可能导致显著的内存泄漏问题。
由于 s0 是包级别变量，只要程序不退出，s0 引用的内存块会一直存活。如果程序长期运行（例如，服务器进程），且 f 被反复调用传入大字符串，内存占用可能持续累积（尽管每次调用 f 只会保留一个内存块的引用）。
*/
func test2(s1 string) {
	s0 = s1[:50]
	// Now, s0 shares the same underlying memory block
	// with s1. Although s1 is not alive now, but s0
	// is still alive, so the memory block they share
	// couldn't be collected, though there are only 50
	// bytes used in the block and all other bytes in
	// the block become unavailable.
}
