package common

import "time"

// 警报抑制时长
const AlertMuteDuration = time.Hour
// const AlertMuteDuration = time.Second * 5

// 监控间隔
const MonitorInterval = time.Minute
// const MonitorInterval = time.Second

// 监控数据长度
const METRICS_NUM = 120