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
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/growingio/growingio-sdk-go/internal/logger"
	"github.com/growingio/growingio-sdk-go/internal/protobuf"
)

type EventBuilder struct {
	EventName    string
	EventTime    int64
	AnonymousId  string
	LoginUserId  string
	LoginUserKey string
	ItemId       string
	ItemKey      string
	Attributes   map[string]interface{}
}

func BuildCustomEvent(builder *EventBuilder) {
	event := &protobuf.EventV3Dto{
		ProjectKey:   AccountId,
		DataSourceId: DataSourceId,
		SdkVersion:   SdkVersion,
		Platform:     Platform,
	}

	event.EventType = protobuf.EventType_CUSTOM
	event.EventName = builder.EventName

	timestamp := builder.getEventTime()
	event.Timestamp = timestamp
	event.SendTime = timestamp

	event.DeviceId = builder.getAnonymousId()
	event.UserId = builder.getLoginUserId()
	event.UserKey = builder.getLoginUserKey()
	event.Attributes = builder.getAttributes()

	sendOrStoreEvent(event)

	logger.Debug("BuildCustomEvent", "event", event.String())
}

func BuildUserLoginEvent(builder *EventBuilder) {
	event := &protobuf.EventV3Dto{
		ProjectKey:   AccountId,
		DataSourceId: DataSourceId,
		SdkVersion:   SdkVersion,
		Platform:     Platform,
	}

	event.EventType = protobuf.EventType_LOGIN_USER_ATTRIBUTES

	timestamp := builder.getEventTime()
	event.Timestamp = timestamp
	event.SendTime = timestamp

	event.DeviceId = builder.getAnonymousId()
	event.UserId = builder.getLoginUserId()
	event.UserKey = builder.getLoginUserKey()
	event.Attributes = builder.getAttributes()

	sendOrStoreEvent(event)

	logger.Debug("BuildUserLoginEvent", "event", event.String())
}

func BuildItemEvent(builder *EventBuilder) {
	event := &protobuf.ItemDto{
		ProjectKey:   AccountId,
		DataSourceId: DataSourceId,
	}

	event.Id = builder.getItemId()
	event.Key = builder.getItemKey()
	event.Attributes = builder.getAttributes()

	sendOrStoreItem(event)

	logger.Debug("BuildItemEvent", "event", event.String())
}

func sendOrStoreEvent(event *protobuf.EventV3Dto) {
	if BatchEnable {
		bInst.pushEvent(event)
	} else {
		sendEvent(event)
	}
}

func sendOrStoreItem(event *protobuf.ItemDto) {
	if BatchEnable {
		bInst.pushItem(event)
	} else {
		sendItem(event)
	}
}

func (e *EventBuilder) getEventTime() int64 {
	if e.EventTime != 0 {
		return e.EventTime
	}
	return time.Now().UnixMilli()
}

func (e *EventBuilder) getAnonymousId() string {
	if e.AnonymousId != "" {
		return e.AnonymousId
	}
	return ""
}

func (e *EventBuilder) getLoginUserId() string {
	if e.LoginUserId != "" {
		return e.LoginUserId
	}
	return ""
}

func (e *EventBuilder) getLoginUserKey() string {
	if e.LoginUserKey != "" {
		return e.LoginUserKey
	}
	return ""
}

func (e *EventBuilder) getItemId() string {
	if e.ItemId != "" {
		return e.ItemId
	}
	return ""
}

func (e *EventBuilder) getItemKey() string {
	if e.ItemKey != "" {
		return e.ItemKey
	}
	return ""
}

func (e *EventBuilder) getAttributes() map[string]string {
	if e.Attributes == nil {
		return make(map[string]string)
	}

	result := make(map[string]string)
	seperator := "||"
	for key, value := range e.Attributes {
		switch v := value.(type) {
		case string:
			result[key] = v
		case []string:
			result[key] = strings.Join(v, seperator)
		case []bool:
			strArr := convertToStringArray(v)
			result[key] = strings.Join(strArr, seperator)
		case []uint8:
			strArr := convertToStringArray(v)
			result[key] = strings.Join(strArr, seperator)
		case []int:
			strArr := convertToStringArray(v)
			result[key] = strings.Join(strArr, seperator)
		case []int32:
			strArr := convertToStringArray(v)
			result[key] = strings.Join(strArr, seperator)
		case []int64:
			strArr := convertToStringArray(v)
			result[key] = strings.Join(strArr, seperator)
		case []float64:
			strArr := convertToStringArray(v)
			result[key] = strings.Join(strArr, seperator)
		case []interface{}:
			strArr := convertToStringArray(v)
			result[key] = strings.Join(strArr, seperator)
		case map[string]string:
			jsonData, _ := json.Marshal(v)
			result[key] = string(jsonData)
		default:
			result[key] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

func convertToStringArray[T any](arr []T) []string {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		strArr[i] = fmt.Sprintf("%v", v)
	}
	return strArr
}
