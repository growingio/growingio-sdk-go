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
	"runtime"
	"time"

	"github.com/growingio/growingio-sdk-go/internal/logger"
	"github.com/growingio/growingio-sdk-go/internal/protobuf"
)

const (
	DefaultMaxSize      = 500
	DefaultFlushAfter   = 10
	DefaultMaxCacheSize = 10240
)

// getDefaultRoutineCount returns an optimal routine count based on system resources
func getDefaultRoutineCount() int {
	// For I/O-intensive work (HTTP requests), use 2x CPU cores
	count := runtime.NumCPU() * 2
	
	// Set reasonable bounds
	if count < 4 {
		count = 4   // Minimum for responsiveness
	}
	if count > 32 {
		count = 32  // Maximum to avoid excessive overhead
	}
	
	return count
}

var (
	BatchEnable  bool
	MaxSize      int
	FlushAfter   int
	RoutineCount int
	MaxCacheSize int

	bInst Batch
)

type Batch interface {
	pushEvent(event *protobuf.EventV3Dto)
	pushItem(item *protobuf.ItemDto)
	pop()
}

type batch struct {
	events chan *protobuf.EventV3Dto
	items  chan *protobuf.ItemDto
}

func InitBatch() {
	if MaxSize == 0 {
		MaxSize = DefaultMaxSize
	}
	if FlushAfter == 0 {
		FlushAfter = DefaultFlushAfter
	}
	if RoutineCount == 0 {
		RoutineCount = getDefaultRoutineCount()
	}
	if MaxCacheSize == 0 {
		MaxCacheSize = DefaultMaxCacheSize
	}

	bInst = NewBatch()
}

func RunBatch() {
	go run(bInst)
}

func NewBatch() Batch {
	return &batch{
		events: make(chan *protobuf.EventV3Dto, MaxCacheSize),
		items:  make(chan *protobuf.ItemDto, MaxCacheSize),
	}
}

func run(b Batch) {
	for i := 0; i < RoutineCount; i++ {
		go func() {
			for {
				b.pop()
			}
		}()
	}
}

func (b *batch) pushEvent(event *protobuf.EventV3Dto) {
	b.events <- event
}

func (b *batch) pushItem(item *protobuf.ItemDto) {
	b.items <- item
}

func (b *batch) pop() {
	var events []*protobuf.EventV3Dto
	var items []*protobuf.ItemDto

Query:
	for {
		select {
		case e := <-b.events:
			events = append(events, e)
			if len(events) >= MaxSize {
				logger.Debug("sending events due to exceeding limit", "count", len(events), "limit", MaxSize)
				sendEvents(events)
				events = make([]*protobuf.EventV3Dto, 0)
			}
		case i := <-b.items:
			items = append(items, i)
			if len(items) >= MaxSize {
				logger.Debug("sending items due to exceeding limit", "count", len(items), "limit", MaxSize)
				sendItems(items)
				items = make([]*protobuf.ItemDto, 0)
			}
		case <-time.After(time.Duration(FlushAfter) * time.Second):
			break Query
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
