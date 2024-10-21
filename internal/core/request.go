package core

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/growingio/growingio-sdk-go/internal/protobuf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	logger "github.com/growingio/growingio-sdk-go/internal/logger"
)

type RequestUrl struct {
	Collect string
	Item    string
}

var Urls RequestUrl
var RequestTimeout int

func sendItem(item *protobuf.ItemDto) {
	items := &protobuf.ItemDtoList{
		Values: []*protobuf.ItemDto{item},
	}
	makeRequest(items, false)
}

func sendItems(items []*protobuf.ItemDto) {
	itemList := &protobuf.ItemDtoList{
		Values: items,
	}
	makeRequest(itemList, false)
}

func sendEvent(event *protobuf.EventV3Dto) {
	eventList := &protobuf.EventV3List{
		Values: []*protobuf.EventV3Dto{event},
	}
	makeRequest(eventList, true)
}

func sendEvents(events []*protobuf.EventV3Dto) {
	eventList := &protobuf.EventV3List{
		Values: events,
	}
	makeRequest(eventList, true)
}

func makeRequest(m protoreflect.ProtoMessage, isEventRequest bool) {
	data, _ := proto.Marshal(m)
	var url string
	if isEventRequest {
		url = Urls.Collect
	} else {
		url = Urls.Item
	}
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	req, _ := http.NewRequest(http.MethodPost, url+"?stm="+timestamp, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/protobuf")
	req.Header.Set("X-Timestamp", timestamp)
	req.Header.Set("Accept", "application/json")

	logger.Debug("make request", "url", req.URL)
	client := &http.Client{
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err, "request failed")
		return
	}
	defer resp.Body.Close()

	logger.Debug("get response", "status", resp.Status)
}
