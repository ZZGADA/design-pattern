package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
	"time"
)

// prometheusMonitorMiddleware count middle ware
func (skw *MetricsCollector) PrometheusMonitorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()

		fmt.Println("-------PrometheusMonitorMiddleware[start]-------", path)

		// 执行
		c.Next()

		fmt.Println("-------PrometheusMonitorMiddleware[ending]-------", path)

		duration := time.Since(start).Seconds()
		statusCode := c.Writer.Status()

		value := &Value{
			Method:  c.Request.Method,
			Path:    c.FullPath(),
			Host:    c.Request.Host,
			Success: statusCode >= 200 && statusCode < 300,
			MType:   MetricType_Http,
			Rt:      time.Duration(duration * float64(time.Second)),
		}

		skw.IncRequestCount(value)
		skw.RequestDuration(value)
	}
}

// MetricsCollector 用于收集和暴露 Prometheus 指标
type MetricsCollector struct {
	requestCount    *prometheus.CounterVec
	requestDuration *prometheus.SummaryVec

	dependRequestDuration *prometheus.SummaryVec
	dependRequestCount    *prometheus.CounterVec

	dbStatusCount *prometheus.GaugeVec
}

func NewMetricsCollector(serviceName string) *MetricsCollector {
	m := new(MetricsCollector)

	//响应时间相关
	// 平均响应时间: rate(rpc_request_duration_sum[5m]) / rate(rpc_request_duration_count[5m])
	m.requestDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:        "request_duration",
			Help:        "request duration",
			ConstLabels: prometheus.Labels{"serverName": serviceName},
			Objectives:  map[float64]float64{0.6: 0.05, 0.9: 0.01, 0.99: 0.001}, //百分位
		},
		//method,rt,type(1:redis 2:mysql 3:kafka 4:rpcx)
		//  Prometheus 中，一个指标（Metric）由它的名称和一组唯一的标签键值对来定义一个时间序列（Time Series）。
		// 绝对不能将持续可变字段（如duration）放在这里，因为它们会导致指标的时间序列数量爆炸。
		// 如果将duration加入进来。这意味着每一次API请求都会创建一个全新的时间序列。
		// 标签的值必须是有限的、可预测的集合
		// 用户ID、请求ID、精确的时间戳、请求耗时
		//  Summary 和 Histogram 类型的指标本身就是用来统计和观察值的分布的（比如请求耗时），我们应该把耗时作为 Observe() 的参数传入，而不是作为标签。
		[]string{"method", "type"},
	)

	m.requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			ConstLabels: prometheus.Labels{"serverName": serviceName},
			Name:        "request_count",
			Help:        "request count",
		},
		//  Prometheus 中，一个指标（Metric）由它的名称和一组唯一的标签键值对来定义一个时间序列（Time Series）。
		[]string{"method", "success", "type", "path"},
	)

	// 注册指标：将创建的指标（如 Counter、Summary、Gauge、Histogram 等）注册到 Prometheus 的默认注册表中，使这些指标能够被 Prometheus 服务器抓取。
	// 错误处理：与 prometheus.Register 不同，MustRegister 会在注册失败时直接触发 panic（而不是返回错误）。这通常用于程序初始化阶段，当指标注册失败时（如指标名称重复），意味着程序配置存在严重问题，应当终止运行。
	// 使用场景：一般在程序启动时，定义完所有需要的指标后调用，确保指标正确注册。

	// 通过 MustRegister 注册的指标，会被包含在默认注册表中，当通过 promhttp.Handler() 暴露指标端点时，这些指标就会被 Prometheus 采集。
	prometheus.MustRegister(m.requestDuration)
	prometheus.MustRegister(m.requestCount)

	return m
}

type MetricType string

const (
	MetricType_Redis MetricType = "redis"
	MetricType_SqlDB MetricType = "sql_db"
	MetricType_Kafka MetricType = "kafka"
	MetricType_Http  MetricType = "http"
	MetricType_Grpc  MetricType = "grpc"
)

type Value struct {
	Host    string
	Path    string
	Method  string
	Success bool
	MType   MetricType
	Rt      time.Duration // 毫秒值
}

func (m *MetricsCollector) IncRequestCount(val *Value) {
	if val == nil {
		return
	}

	// WithLabelValues 创建指标
	method := strings.Trim(val.Method, "*")
	m.requestCount.WithLabelValues(method, val.Path, fmt.Sprintf("%t", val.Success), string(val.MType)).Inc()
}

func (m *MetricsCollector) RequestDuration(val *Value) {
	if val == nil {
		return
	}
	rt := val.Rt / time.Millisecond
	method := strings.Trim(val.Method, "*")

	// Observe 函数用于记录一次请求的耗时（响应时间），它会将传入的数值（如毫秒）加入到 SummaryVec 指标的统计中。Prometheus 会自动根据所有 Observe 的值计算平均值、分位数等统计信息。因此，可以用它来统计接口的平均响应时间、P99 等性能指标
	m.requestDuration.WithLabelValues(method, string(val.MType)).Observe(float64(rt))
}
