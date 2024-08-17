package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/117503445/goutils"
)

func push(name string, hour int) {
	fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
	httpClient := &http.Client{}
	url := fmt.Sprintf("https://push.gh.117503445.top:20000/push/text/v1?name=%s", name)
	fmt.Println(url, hour)
	r, err := httpClient.Post(url, "application/json", strings.NewReader(fmt.Sprintf("电脑已开机%d小时", hour)))
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}


func a() {

	goutils.InitZeroLog()

	fmt.Println(a)
	// 获取初始的网络接口统计数据
	initialStats, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("Error getting initial network stats:", err)
		return
	}

	// 等待一段时间
	time.Sleep(1 * time.Second)

	// 获取新的网络接口统计数据
	newStats, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("Error getting new network stats:", err)
		return
	}

	// 计算网络速度
	for i, initialStat := range initialStats {
		newStat := newStats[i]
		bytesSentPerSec := newStat.BytesSent - initialStat.BytesSent
		bytesRecvPerSec := newStat.BytesRecv - initialStat.BytesRecv
		if bytesSentPerSec > 2000 || bytesRecvPerSec > 2000 {
			fmt.Println("Network Interface: ", initialStat.Name)
			fmt.Println("Bytes Sent/sec: ", bytesSentPerSec)
			fmt.Println("Bytes Recv/sec: ", bytesRecvPerSec)
		}
	}
}

func main() {

	go func() {
		for {
			a()
			// time.Sleep(5 * time.Second)
		}
	}()

	var err error
	var idleTime time.Duration

	oneSecond, _ := time.ParseDuration("1s")

	for err == nil {
		

		if idleTime.Seconds() > 1.0 {
			// log.Printf("Idle for %d seconds.", int(idleTime.Seconds()))
		}

		time.Sleep(oneSecond)
	}

	if err != nil {
		log.Fatal(err)
	}

}
