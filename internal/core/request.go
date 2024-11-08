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
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/growingio/growingio-sdk-go/internal/protobuf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	logger "github.com/growingio/growingio-sdk-go/internal/logger"
)

type Request struct {
	URL       string
	Headers   map[string]string
	Body      []byte
	Timestamp int64
}

type RequestUrl struct {
	Collect string
	Item    string
}

var Urls RequestUrl
var RequestTimeout int = 5

func sendItem(item *protobuf.ItemDto) {
	itemList := &protobuf.ItemDtoList{
		Values: []*protobuf.ItemDto{item},
	}
	makeRequest(itemList, Urls.Item)
}

func sendItems(items []*protobuf.ItemDto) {
	itemList := &protobuf.ItemDtoList{
		Values: items,
	}
	makeRequest(itemList, Urls.Item)
}

func sendEvent(event *protobuf.EventV3Dto) {
	eventList := &protobuf.EventV3List{
		Values: []*protobuf.EventV3Dto{event},
	}
	makeRequest(eventList, Urls.Collect)
}

func sendEvents(events []*protobuf.EventV3Dto) {
	eventList := &protobuf.EventV3List{
		Values: events,
	}
	makeRequest(eventList, Urls.Collect)
}

func makeRequest(m protoreflect.ProtoMessage, baseURL string) {
	timestamp := time.Now().UnixMilli()
	timestampString := fmt.Sprintf("%d", timestamp)
	url := baseURL + "?stm=" + timestampString
	headers := make(map[string]string)
	headers["Content-Type"] = "application/protobuf"
	headers["Accept"] = "application/json"
	headers["X-Timestamp"] = timestampString
	body, _ := proto.Marshal(m)

	req := &Request{
		URL:       url,
		Headers:   headers,
		Body:      body,
		Timestamp: timestamp,
	}

	pm := getPipeManager()
	if err := pm.execute(req); err != nil {
		logger.Error(err, "make request failed")
		return
	}

	if err := sendRequest(req); err != nil {
		logger.Error(err, "send request failed")
	}
}

func sendRequest(req *Request) error {
	httpReq, _ := http.NewRequest(http.MethodPost, req.URL, bytes.NewBuffer(req.Body))
	logger.Debug("make request", "url", httpReq.URL)

	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	logger.Debug("get response", "status", resp.Status)
	return nil
}
