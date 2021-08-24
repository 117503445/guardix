package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

func main() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	name := viper.GetString("name")
	hour := 0
	push(name, hour)

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
				push(name, hour)
				println("end 1 hour")
			} else if maxTime == 3600*2 {
				maxTime = 3600 * 4
				startTime = time.Now()
				hour += 2
				push(name, hour)
				println("end 2 hour")
			} else if maxTime == 3600*4 {
				startTime = time.Now()
				hour += 4
				push(name, hour)
				println("end 4 hour")
			}
		}

	}
}
