package core

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/snappy"
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
var RequestTimeout int

func sendItem(item *protobuf.ItemDto) {
	items := &protobuf.ItemDtoList{
		Values: []*protobuf.ItemDto{item},
	}
	makeRequest(items, Urls.Item)
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

type Pipe func(*Request) error

type PipeManager struct {
	pipes []Pipe
}

var pipe *PipeManager

func GetPipeManager() *PipeManager {
	if pipe == nil {
		pipe = &PipeManager{}
		pipe.AddPipe(compress)
		pipe.AddPipe(encrypt)
	}
	return pipe
}

func (pm *PipeManager) AddPipe(pipe Pipe) {
	pm.pipes = append(pm.pipes, pipe)
}

func (pm *PipeManager) Execute(req *Request) error {
	for _, pipe := range pm.pipes {
		if err := pipe(req); err != nil {
			return err
		}
	}
	return nil
}

func compress(req *Request) error {
	compressed := snappy.Encode(nil, req.Body)
	req.Body = compressed
	req.Headers["X-Compress-Codec"] = "2"
	return nil
}

func encrypt(req *Request) error {
	hint := byte(req.Timestamp % 256)
	encrypted := make([]byte, len(req.Body))
	for i, b := range req.Body {
		encrypted[i] = b ^ hint
	}
	req.Body = encrypted
	req.Headers["X-Crypt-Codec"] = "1"
	return nil
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

	pm := GetPipeManager()
	if err := pm.Execute(req); err != nil {
		logger.Error(err, "make request failed")
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
