package main

import "bridge"

func main() {
	cpuIntel := bridge.NewIntelCPU()
	cpuAmd := bridge.NewAmdCPU()

	randomIntel := bridge.NewRandomIORefined(cpuIntel)
	randomAmd := bridge.NewRandomIORefined(cpuAmd)
	seqIntel := bridge.NewSequentialIORefined(cpuIntel)
	seqAmd := bridge.NewSequentialIORefined(cpuAmd)

	slice := make([]bridge.IOAbstraction, 4, 4)
	slice[0] = randomIntel
	slice[1] = randomAmd
	slice[2] = seqIntel
	slice[3] = seqAmd

	for i := 0; i < 4; i++ {
		slice[i].IO("执行")
	}

}
