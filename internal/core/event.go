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

func (b EventBuilder) BuildCustomEvent() {
	event := &protobuf.EventV3Dto{
		ProjectKey:   AccountId,
		DataSourceId: DataSourceId,
		SdkVersion:   SdkVersion,
		Platform:     Platform,
	}

	event.EventType = protobuf.EventType_CUSTOM
	event.EventName = b.EventName

	timestamp := b.getEventTime()
	event.Timestamp = timestamp
	event.SendTime = timestamp

	event.DeviceId = b.getAnonymousId()
	event.UserId = b.getLoginUserId()
	event.UserKey = b.getLoginUserKey()
	event.Attributes = b.getAttributes()

	if BatchEnable {
		batch.pushEvent(event)
	} else {
		sendEvent(event)
	}

	logger.Debug("BuildCustomEvent", "event", event.String())
}

func (b EventBuilder) BuildUserLoginEvent() {
	event := &protobuf.EventV3Dto{
		ProjectKey:   AccountId,
		DataSourceId: DataSourceId,
		SdkVersion:   SdkVersion,
		Platform:     Platform,
	}

	event.EventType = protobuf.EventType_LOGIN_USER_ATTRIBUTES

	timestamp := b.getEventTime()
	event.Timestamp = timestamp
	event.SendTime = timestamp

	event.DeviceId = b.getAnonymousId()
	event.UserId = b.getLoginUserId()
	event.UserKey = b.getLoginUserKey()
	event.Attributes = b.getAttributes()

	if BatchEnable {
		batch.pushEvent(event)
	} else {
		sendEvent(event)
	}

	logger.Debug("BuildUserLoginEvent", "event", event.String())
}

func (b EventBuilder) BuildItemEvent() {
	event := &protobuf.ItemDto{
		ProjectKey:   AccountId,
		DataSourceId: DataSourceId,
	}

	event.Id = b.getItemId()
	event.Key = b.getItemKey()
	event.Attributes = b.getAttributes()

	if BatchEnable {
		batch.pushItem(event)
	} else {
		sendItem(event)
	}

	logger.Debug("BuildItemEvent", "event", event.String())
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
