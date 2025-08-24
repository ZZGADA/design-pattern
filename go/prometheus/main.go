package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包
	"runtime"
	"sync"
	"time"

	"github.ZZGADA.com/prometheus/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	e := gin.New()

	initPrometheus(e)
	initPprof(e)
	initMemoryLeakRoutes(e)

	// 启动内存泄漏场景
	//startMemoryLeakScenarios()

	log.Info("服务器启动成功，端口: 8088")
	log.Info("Prometheus metrics: http://localhost:8088/metrics")
	log.Info("Pprof debug: http://localhost:8088/debug/pprof/")
	log.Info("内存泄漏触发: http://localhost:8088/leak/")

	if err := e.Run(":8088"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

func initPrometheus(e *gin.Engine) {
	// Register Prometheus metrics handler
	// This is a placeholder for the actual implementation
	// In a real application, you would set up the Prometheus metrics endpoint here
	// e.g., http.Handle("/metrics", promhttp.Handler())
	log.Info("Prometheus metrics handler registered")

	collector := middleware.NewMetricsCollector("class_svc")
	use := e.Use(collector.PrometheusMonitorMiddleware())

	// Register the Prometheus metrics endpoint
	// 注册端点
	//    1. Default Collectors: The github.com/prometheus/client_golang library comes with a set of built-in collectors for standard Go application metrics. These are automatically registered when you use the default
	//      promhttp.Handler.
	//
	//
	//   2. `promhttp.Handler()`: In your initPrometheus function, this line is the key:
	//
	//
	//   1     use.GET("/metrics", gin.WrapH(promhttp.Handler()))
	//
	//      The promhttp.Handler() serves metrics from a default global registry. This default registry automatically includes the Go runtime collector.
	//
	//
	//   3. Go Runtime Collector: This built-in collector queries the Go runtime for various performance statistics, including:
	//       * Garbage collection stats (which produces go_gc_duration_seconds).
	//       * Number of goroutines (go_goroutines).
	//       * Memory statistics (go_memstats_* metrics).
	//       * Number of OS threads (go_threads).
	use.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

// 内存泄漏相关的全局变量
var (
	// slice 持续增长场景
	globalSlice [][]byte
	sliceMutex  sync.RWMutex

	// map 持续增长场景
	globalMap = make(map[string][]byte)
	mapMutex  sync.RWMutex

	// goroutine 泄漏场景
	leakedChannels []chan struct{}
	channelMutex   sync.Mutex

	// 对象缓存场景
	objectCache = make(map[string]*LargeObject)
	cacheMutex  sync.RWMutex
)

// 大对象结构体
type LargeObject struct {
	Data      []byte
	Timestamp time.Time
	Metadata  map[string]interface{}
}

// 初始化 pprof
func initPprof(e *gin.Engine) {
	// 添加 pprof 路由组
	pprof.Register(e, "/debug/pprof")

	// 添加内存状态查看端点
	e.GET("/debug/memory-status", getMemoryStatus)

	log.Info("Pprof debug endpoints initialized")
}

// 初始化内存泄漏路由
func initMemoryLeakRoutes(e *gin.Engine) {
	leakGroup := e.Group("/leak")
	{
		leakGroup.GET("/slice", triggerSliceLeak)
		leakGroup.GET("/map", triggerMapLeak)
		leakGroup.GET("/goroutine", triggerGoroutineLeak)
		leakGroup.GET("/object-cache", triggerObjectCacheLeak)
		leakGroup.GET("/force-gc", forceGC)
		leakGroup.GET("/status", getLeakStatus)
		leakGroup.GET("/clear", clearLeaks)
	}

	log.Info("Memory leak routes initialized")
}

// 启动内存泄漏场景
func startMemoryLeakScenarios() {
	// 启动定时的内存泄漏场景
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			// 缓慢的 slice 增长
			slowSliceGrowth()

			// 缓慢的 map 增长
			slowMapGrowth()
		}
	}()

	log.Info("Memory leak scenarios started")
}

// 缓慢的 slice 增长
func slowSliceGrowth() {
	sliceMutex.Lock()
	defer sliceMutex.Unlock()

	// 每次添加一小块内存
	data := make([]byte, 1024*10) // 10KB
	for i := range data {
		data[i] = byte(i % 256)
	}
	globalSlice = append(globalSlice, data)

	// 保持一定数量的切片，但不清理旧的
	if len(globalSlice) > 1000 {
		// 只清理一小部分，造成缓慢泄漏
		globalSlice = globalSlice[10:]
	}
}

// 缓慢的 map 增长
func slowMapGrowth() {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	key := fmt.Sprintf("key_%d", time.Now().UnixNano())
	data := make([]byte, 1024*5) // 5KB
	for i := range data {
		data[i] = byte(i % 256)
	}
	globalMap[key] = data

	// 保持 map 大小，但不完全清理
	if len(globalMap) > 500 {
		// 只删除少量 key，造成泄漏
		count := 0
		for k := range globalMap {
			delete(globalMap, k)
			count++
			if count >= 5 {
				break
			}
		}
	}
}

// 触发 slice 泄漏
func triggerSliceLeak(c *gin.Context) {
	count := 100 // 默认创建100个slice

	sliceMutex.Lock()
	defer sliceMutex.Unlock()

	for i := 0; i < count; i++ {
		// 创建大的 slice
		data := make([]byte, 1024*100) // 100KB
		for j := range data {
			data[j] = byte(j % 256)
		}
		globalSlice = append(globalSlice, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      fmt.Sprintf("Created %d slices, total slices: %d", count, len(globalSlice)),
		"memory_usage": fmt.Sprintf("~%d MB", len(globalSlice)*100/1024),
	})
}

// 触发 map 泄漏
func triggerMapLeak(c *gin.Context) {
	count := 100

	mapMutex.Lock()
	defer mapMutex.Unlock()

	for i := 0; i < count; i++ {
		key := fmt.Sprintf("leak_key_%d_%d", time.Now().UnixNano(), i)
		data := make([]byte, 1024*50) // 50KB
		for j := range data {
			data[j] = byte(j % 256)
		}
		globalMap[key] = data
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      fmt.Sprintf("Created %d map entries, total entries: %d", count, len(globalMap)),
		"memory_usage": fmt.Sprintf("~%d MB", len(globalMap)*50/1024),
	})
}

// 触发 goroutine 泄漏
func triggerGoroutineLeak(c *gin.Context) {
	count := 50

	channelMutex.Lock()
	defer channelMutex.Unlock()

	for i := 0; i < count; i++ {
		ch := make(chan struct{})
		leakedChannels = append(leakedChannels, ch)

		// 创建永远不会结束的 goroutine
		go func(id int, ch chan struct{}) {
			data := make([]byte, 1024*10) // 10KB per goroutine
			for j := range data {
				data[j] = byte(j % 256)
			}

			// 永远等待，不会退出
			select {
			case <-ch:
				return
			case <-time.After(time.Hour * 24 * 365): // 等待一年
				return
			}
		}(i, ch)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         fmt.Sprintf("Created %d leaked goroutines, total goroutines: %d", count, runtime.NumGoroutine()),
		"leaked_channels": len(leakedChannels),
	})
}

// 触发对象缓存泄漏
func triggerObjectCacheLeak(c *gin.Context) {
	count := 50

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	for i := 0; i < count; i++ {
		key := fmt.Sprintf("cache_obj_%d_%d", time.Now().UnixNano(), i)

		// 创建大对象
		obj := &LargeObject{
			Data:      make([]byte, 1024*200), // 200KB
			Timestamp: time.Now(),
			Metadata: map[string]interface{}{
				"id":      i,
				"created": time.Now(),
				"type":    "leaked_object",
				"size":    1024 * 200,
			},
		}

		// 填充数据
		for j := range obj.Data {
			obj.Data[j] = byte(j % 256)
		}

		objectCache[key] = obj
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      fmt.Sprintf("Created %d cached objects, total objects: %d", count, len(objectCache)),
		"memory_usage": fmt.Sprintf("~%d MB", len(objectCache)*200/1024),
	})
}

// 强制执行垃圾回收
func forceGC(c *gin.Context) {
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	runtime.GC()
	runtime.GC() // 执行两次确保彻底回收

	runtime.ReadMemStats(&m2)

	c.JSON(http.StatusOK, gin.H{
		"message": "Forced garbage collection completed",
		"before_gc": gin.H{
			"alloc_mb":       m1.Alloc / 1024 / 1024,
			"total_alloc_mb": m1.TotalAlloc / 1024 / 1024,
			"sys_mb":         m1.Sys / 1024 / 1024,
			"num_gc":         m1.NumGC,
		},
		"after_gc": gin.H{
			"alloc_mb":       m2.Alloc / 1024 / 1024,
			"total_alloc_mb": m2.TotalAlloc / 1024 / 1024,
			"sys_mb":         m2.Sys / 1024 / 1024,
			"num_gc":         m2.NumGC,
		},
		"freed_mb": (m1.Alloc - m2.Alloc) / 1024 / 1024,
	})
}

// 获取泄漏状态
func getLeakStatus(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	sliceMutex.RLock()
	sliceCount := len(globalSlice)
	sliceMutex.RUnlock()

	mapMutex.RLock()
	mapCount := len(globalMap)
	mapMutex.RUnlock()

	channelMutex.Lock()
	channelCount := len(leakedChannels)
	channelMutex.Unlock()

	cacheMutex.RLock()
	cacheCount := len(objectCache)
	cacheMutex.RUnlock()

	c.JSON(http.StatusOK, gin.H{
		"memory_stats": gin.H{
			"alloc_mb":       m.Alloc / 1024 / 1024,
			"total_alloc_mb": m.TotalAlloc / 1024 / 1024,
			"sys_mb":         m.Sys / 1024 / 1024,
			"heap_alloc_mb":  m.HeapAlloc / 1024 / 1024,
			"heap_sys_mb":    m.HeapSys / 1024 / 1024,
			"num_gc":         m.NumGC,
			"num_goroutine":  runtime.NumGoroutine(),
		},
		"leak_status": gin.H{
			"global_slices":      sliceCount,
			"global_map_keys":    mapCount,
			"leaked_channels":    channelCount,
			"cached_objects":     cacheCount,
			"estimated_slice_mb": sliceCount * 100 / 1024,
			"estimated_map_mb":   mapCount * 50 / 1024,
			"estimated_cache_mb": cacheCount * 200 / 1024,
		},
	})
}

// 清理泄漏（用于测试）
func clearLeaks(c *gin.Context) {
	// 清理 slice
	sliceMutex.Lock()
	globalSlice = nil
	sliceMutex.Unlock()

	// 清理 map
	mapMutex.Lock()
	for k := range globalMap {
		delete(globalMap, k)
	}
	mapMutex.Unlock()

	// 关闭泄漏的 goroutines
	channelMutex.Lock()
	for _, ch := range leakedChannels {
		close(ch)
	}
	leakedChannels = nil
	channelMutex.Unlock()

	// 清理缓存
	cacheMutex.Lock()
	for k := range objectCache {
		delete(objectCache, k)
	}
	cacheMutex.Unlock()

	// 强制垃圾回收
	runtime.GC()
	runtime.GC()

	c.JSON(http.StatusOK, gin.H{
		"message": "All leaks cleared and garbage collection forced",
	})
}

// 获取内存状态
func getMemoryStatus(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.JSON(http.StatusOK, gin.H{
		"memory_stats": gin.H{
			"alloc_mb":         m.Alloc / 1024 / 1024,        // 当前分配的内存
			"total_alloc_mb":   m.TotalAlloc / 1024 / 1024,   // 总分配的内存
			"sys_mb":           m.Sys / 1024 / 1024,          // 系统内存
			"heap_alloc_mb":    m.HeapAlloc / 1024 / 1024,    // 堆分配的内存
			"heap_sys_mb":      m.HeapSys / 1024 / 1024,      // 堆系统内存
			"heap_idle_mb":     m.HeapIdle / 1024 / 1024,     // 堆空闲内存
			"heap_inuse_mb":    m.HeapInuse / 1024 / 1024,    // 堆使用中的内存
			"heap_released_mb": m.HeapReleased / 1024 / 1024, // 已释放给OS的内存
			"num_gc":           m.NumGC,                      // GC次数
			"gc_cpu_fraction":  m.GCCPUFraction,              // GC消耗的CPU时间比例
			"num_goroutine":    runtime.NumGoroutine(),       // goroutine数量
		},
		"gc_stats": gin.H{
			"last_gc":        time.Unix(0, int64(m.LastGC)).Format("2006-01-02 15:04:05"),
			"pause_total_ns": m.PauseTotalNs,
			"pause_ns":       m.PauseNs,
		},
		"instructions": []string{
			"使用 /debug/pprof/heap 查看堆内存详情",
			"使用 /debug/pprof/goroutine 查看goroutine详情",
			"使用 /debug/pprof/profile 进行CPU性能分析",
			"使用 /leak/ 下的接口触发内存泄漏场景",
		},
	})
}
