//
// @license
// Copyright (C) 2024 Beijing Yishu Technology Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"sync"
	"testing"
	"time"

	"github.com/go-logr/stdr"
	"github.com/growingio/growingio-sdk-go/internal/logger"
	"github.com/growingio/growingio-sdk-go/internal/protobuf"
)

func createEvent(seqId int32) *protobuf.EventV3Dto {
	event := &protobuf.EventV3Dto{
		ProjectKey:   "123456",
		DataSourceId: "654321",
		SdkVersion:   "1.0.0",
		Platform:     "go",
	}

	event.EventType = protobuf.EventType_CUSTOM
	event.EventName = "name_" + string(seqId)

	timestamp := time.Now().UnixMilli()
	event.Timestamp = timestamp
	event.SendTime = timestamp

	event.UserId = "user1"
	return event
}

func TestSendEvent(t *testing.T) {
	// 手动修改以下几个参数，进行测试
	// 是否batch发送
	BatchEnable = false
	// 单包大小
	MaxSize = 500
	// 设置goroutine数
	RoutineCount = 3
	// 设置超时间隔
	FlushAfter = 3
	// 模拟事件数量
	eventCount := int32(MaxSize*2 + (MaxSize - 1))
	// 内存中缓存的最大事件数量，一般不修改
	MaxCacheSize = 10240

	// 打开日志
	stdr.SetVerbosity(8)
	logger.Debug("begin TestSendEvent")

	// 初始化batch
	InitBatch()
	// 模拟埋点事件
	for i := int32(0); i < eventCount; i++ {
		go batch.pushEvent(createEvent(i))
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
	}()
	wg.Wait()
}
