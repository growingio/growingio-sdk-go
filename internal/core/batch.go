package core

import (
	"time"

	"github.com/growingio/growingio-sdk-go/internal/logger"
	"github.com/growingio/growingio-sdk-go/internal/protobuf"
)

const defaultMaxSize = 500
const defaultFlushAfter = 10
const defaultRoutineCount = 16
const defaultMaxCacheSize = 10240

var (
	BatchEnable bool
	MaxSize     int
	FlushAfter  int

	RoutineCount int
	MaxCacheSize int

	batch *Batch
)

type Batch struct {
	routine chan struct{}
	events  chan *protobuf.EventV3Dto
	items   chan *protobuf.ItemDto
}

func InitBatch() {
	if MaxSize == 0 {
		MaxSize = defaultMaxSize
	}
	if FlushAfter == 0 {
		FlushAfter = defaultFlushAfter
	}
	if RoutineCount == 0 {
		RoutineCount = defaultRoutineCount
	}
	if MaxCacheSize == 0 {
		MaxCacheSize = defaultMaxCacheSize
	}

	batch = &Batch{
		events:  make(chan *protobuf.EventV3Dto, MaxCacheSize),
		items:   make(chan *protobuf.ItemDto, MaxCacheSize),
		routine: make(chan struct{}, RoutineCount),
	}

	go send()
}

func send() {
	for {
		batch.routine <- struct{}{}
		go func() {
			defer func() { <-batch.routine }()
			batch.pop()
		}()
	}
}

func (batch *Batch) pushEvent(event *protobuf.EventV3Dto) {
	batch.events <- event
}

func (batch *Batch) pushItem(item *protobuf.ItemDto) {
	batch.items <- item
}

func (batch *Batch) pop() {
	var events []*protobuf.EventV3Dto
	var items []*protobuf.ItemDto

L:
	for {
		select {
		case e := <-batch.events:
			events = append(events, e)
			if len(events) >= MaxSize {
				logger.Debug("sending events due to exceeding limit", "count", len(events), "limit", MaxSize)
				sendEvents(events)
				events = make([]*protobuf.EventV3Dto, 0)
			}
		case i := <-batch.items:
			items = append(items, i)
			if len(items) >= MaxSize {
				logger.Debug("sending items due to exceeding limit", "count", len(items), "limit", MaxSize)
				sendItems(items)
				items = make([]*protobuf.ItemDto, 0)
			}
		case <-time.After(time.Duration(FlushAfter) * time.Second):
			break L
		}
	}

	if len(events) > 0 {
		logger.Debug("sending events due to timeout", "count", len(events), "timeout(s)", FlushAfter)
		sendEvents(events)
	}

	if len(items) > 0 {
		logger.Debug("sending items due to timeout", "count", len(items), "timeout(s)", FlushAfter)
		sendItems(items)
	}
}
