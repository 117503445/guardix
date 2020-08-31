package main

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func serverChanGet(serverChanKey string, hour int) {
	httpClient := &http.Client{}
	_, err := httpClient.Get(fmt.Sprintf("https://sc.ftqq.com/%s.send?text=电脑已开机%d小时", serverChanKey, hour))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	serverChanKey := viper.GetString("server_chan")
	hour := 0
	startTime := time.Now()
	maxTime := 3600 // 1 hour
	for true {
		time.Sleep(time.Second)
		duration := int(time.Now().Sub(startTime))
		durationSecond := duration / 1000000000
		// fmt.Println(durationSecond)
		if durationSecond > maxTime {
			if maxTime == 3600 {
				maxTime = 3600 * 2
				startTime = time.Now()
				hour++
				serverChanGet(serverChanKey, hour)
				println("end 1 hour")
			} else if maxTime == 3600*2 {
				maxTime = 3600 * 4
				startTime = time.Now()
				hour += 2
				serverChanGet(serverChanKey, hour)
				println("end 2 hour")
			} else if maxTime == 3600*4 {
				startTime = time.Now()
				hour += 4
				serverChanGet(serverChanKey, hour)
				println("end 4 hour")
			}
		}

	}
}
