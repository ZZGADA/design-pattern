package utils

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

// GetGoroutineID 打印当前的协程
func GetGoroutineID() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := bytes.Fields(buf[:n])[1]
	id, err := strconv.ParseUint(string(idField), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
