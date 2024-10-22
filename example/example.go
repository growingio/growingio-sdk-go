package main

import (
	"sync"
	"time"

	sdk "github.com/growingio/growingio-sdk-go"
)

func main() {
	// 通过创建Config对象初始化
	initAnalytics()

	// 通过本地配置文件生成Config对象初始化
	// initAnalyticsByConfigFile()

	// 埋点事件
	trackCustomEvent()
	trackCustomEventByChaining()
	trackCustomEventWithLargeAttributes()

	// 用户登录属性事件
	trackUser()

	// 维度表
	submitItem()

	// examplet调试时，运行完会直接退出，增加延时，等待sdk中的异步实现（发送埋点数据）执行完成
	waitGoroutineInSdk()
}

func initAnalytics() {
	config := &sdk.Config{
		SdkConfig: sdk.SdkConfig{
			AccountId:    "0a1b4118dd954ec3bcc69da5138bdb96",
			DataSourceId: "ab555003531e0fd1",
			LogLevel:     sdk.DebugLogLevel,
		},
		HttpConfig: sdk.HttpConfig{
			ServerHost:     "https://napi.growingio.com/",
			RequestTimeout: 5,
		},
		BatchConfig: sdk.BatchConfig{
			Enable:     true,
			MaxSize:    100,
			FlushAfter: 5,
		},
	}
	sdk.InitAnalytics(config)
}

func initAnalyticsByConfigFile() {
	sdk.InitAnalyticsByConfigFile("example/yml/config.yml")
}

func trackCustomEvent() {
	builder := sdk.NewCustomEvent("launch")
	builder.AnonymousId = "70C4B9C6-7B9B-675A-5E6B-4D823F5696E3"
	builder.Attributes = map[string]interface{}{
		"name":   "richMan",
		"age":    100,
		"isRich": true,
		"hobby": []string{
			"football",
			"pingpong",
		},
	}
	sdk.TrackCustomEvent(builder)
}

func trackCustomEventWithLargeAttributes() {
	builder := sdk.NewCustomEvent("launch")
	builder.AnonymousId = "70C4B9C6-7B9B-675A-5E6B-4D823F5696E3"
	builder.Attributes = largeAttributes()
	sdk.TrackCustomEvent(builder)
}

func trackCustomEventByChaining() {
	sdk.TrackCustomEvent(
		sdk.NewCustomEvent("userInfo").WithLoginUserId("david").WithAttributes(map[string]interface{}{
			"name":   "richMan",
			"age":    100,
			"isRich": true,
			"hobby": []string{
				"football",
				"pingpong",
			},
		}))
}

func trackUser() {
	sdk.TrackUser(
		sdk.NewUser("jack").WithEventTime(time.Now().UnixMilli()).WithAttributes(map[string]interface{}{
			"key": "value",
		}))
}

func submitItem() {
	sdk.SubmitItem(
		sdk.NewItem("num100", "apple").WithAttributes(map[string]interface{}{
			"key": "value",
		}))
}

func largeAttributes() map[string]interface{} {
	return map[string]interface{}{
		"key":  "value",
		"key2": true,
		"key3": 100,
		"key4": 23.45,
		"key5": []string{
			"good", "bad",
		},
		"key6": []int{
			1, 2, 3, 4,
		},
		"key7": []float64{
			1.11, 2.22, 3.33, 4.44,
		},
		"key8": map[string]string{
			"subKey": "subValue",
		},
		"key9": map[string]interface{}{
			"subKey2": sdk.Config{},
		},
		"key10": map[string]interface{}{
			"subKey3": map[string]interface{}{
				"subKey4": "subValue4",
			},
		},
		"key11": nil,
	}
}

func waitGoroutineInSdk() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(20 * time.Second)
	}()
	wg.Wait()
}
