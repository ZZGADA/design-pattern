package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	id  int
	msg string
}

type Result struct {
	id  int
	msg string
}

type RoutinePool struct {
	workerNum  int
	taskChan   chan Task
	resultChan chan Result
	success    atomic.Int32

	wg sync.WaitGroup
}

func (pool *RoutinePool) Start() {
	for i := 0; i < pool.workerNum; i++ {
		pool.wg.Add(1)
		go pool.worker()
	}
}

// Listen 监听处理结果
func (pool *RoutinePool) Listen() {
	go func() {
		for result := range pool.resultChan {
			fmt.Println(result.msg)
			pool.success.Add(1)
		}
	}()

	// 阻塞监听状态 等待所有worker 消费结束
	pool.wg.Wait()
}

func (pool *RoutinePool) Close() {
	// close 关闭 resultChan 和 Listen方法中的goroutine
	close(pool.resultChan)
}

// push Task任务等待被消费
func (pool *RoutinePool) push(task Task) {
	// 将任务推入任务队列等待消费
	// 串行向100个工作者发送消息
	// 但100个工作者 并发处理
	pool.taskChan <- task
}

// worker 协程池中的消费者
func (pool *RoutinePool) worker() {
	defer pool.wg.Done()

	// 阻塞从任务队列中获取任务
	for task := range pool.taskChan {
		result := pool.process(&task)
		pool.resultChan <- result
	}

	// close( pool.taskChan) 通知所有worker goroutine结束消费处理（结束信号）
	fmt.Println("pool worker exit")
}

// process 处理逻辑函数
func (pool *RoutinePool) process(task *Task) Result {
	// 模拟处理过程
	time.Sleep(time.Second * 2)
	return Result{
		id:  task.id,
		msg: "task has been processed , task id is " + strconv.Itoa(task.id),
	}
}

// main 主函数
// 一个任务队列（task chan）将消息下发给多个工作者
// 工作者们将对任务进行处理 最终返回给 结果队列（result chan）
func main() {
	// 初始化任务队列中的任务数量（起模拟作用 ）
	// 初始化 worker数量
	const (
		taskNum   = 500
		workerNum = 100
	)

	// 初始化 待消费任务队列 结果队列
	taskChan := make(chan Task)
	resultChan := make(chan Result)

	// 启动协程池
	routinePool := &RoutinePool{
		workerNum:  workerNum,
		taskChan:   taskChan,
		resultChan: resultChan,
	}
	routinePool.Start()

	start := time.Now()
	// 启动任务队列协程
	go func(pool *RoutinePool) {
		// 模拟任务创建 并将任务推送至协程池等待被处理
		for i := 0; i < taskNum; i++ {
			task := Task{
				id:  i,
				msg: fmt.Sprintf("task %d", i),
			}
			pool.push(task)
		}

		// 结束信号
		close(pool.taskChan)
	}(routinePool)

	// 启动结果队列 监听结果
	routinePool.Listen()

	end := time.Now()
	fmt.Printf("successful result is %d and process time is %s \n", routinePool.success.Load(), end.Sub(start).String())

	// 协程池结束
	routinePool.Close()

	// 结论：
	// 1. 使用协程池并发处理1000个任务。从串行执行1000个任务花费 1000 * 2 = 2000s 降至 1000/100 * 2 = 20s
	// 2. push 串行发送 worker并发处理 listen 串行接接收。
	// 3. push 串行依次向taskChan 发送消息消息 同时有100个worker goroutine在阻塞等待接收消息。
	//    因为消费chan的阻塞性 其实100个worker 是串行接收消息 但是是并发处理结果 ==> 所以可以理解为2s批量消费 100个任务
	// 4. result chan 阻塞接收消息 串行打印
	// 5. successful result is 1000 and process time is 20.013605375s   （还有发送消息 和 开辟栈空间等时间）

}
