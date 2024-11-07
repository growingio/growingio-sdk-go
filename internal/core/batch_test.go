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
	"testing"
	"time"

	"github.com/growingio/growingio-sdk-go/internal/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestInitBatch(t *testing.T) {
	InitBatch()

	assert.NotNil(t, bInst, "Batch instance should be initialized")
	assert.NotNil(t, routine, "Routine channel should be initialized")
}

func TestNewBatch(t *testing.T) {
	b := NewBatch()

	assert.NotNil(t, b)
	assert.NotNil(t, b.(*batch).events)
	assert.NotNil(t, b.(*batch).items)

	assert.Equal(t, cap(b.(*batch).events), MaxCacheSize)
	assert.Equal(t, cap(b.(*batch).items), MaxCacheSize)
}

func TestRun(t *testing.T) {
	routineCount := 4
	routine = make(chan struct{}, routineCount)
	mockBatch := NewMockBatch(t)

	// Create a channel to record the current number of pop() callers
	// The buffer size is set to routineCount*2 to avoid blocking
	curCallChan := make(chan struct{}, routineCount*2)

	mockBatch.EXPECT().pop().Run(func() {
		// When pop() is called, send a struct{} to curCallChan
		curCallChan <- struct{}{}
		// Get the current number of pop() callers
		curCount := len(curCallChan)
		// Ensure that the current caller count does not exceed routineCount
		assert.LessOrEqual(t, curCount, routineCount)
		// Read from curCallChan to reduce the caller count after execution
		<-curCallChan
	}).After(10 * time.Microsecond) // Delay pop() execution by 10 microseconds to simulate processing time

	// Start the run() method, which will invoke pop() in a loop
	go run(mockBatch)

	// Wait for 5 seconds to allow run() enough time to make multiple pop() calls
	time.Sleep(5000 * time.Millisecond)
	mockBatch.AssertExpectations(t)
}

func TestPushEvent(t *testing.T) {
	b := NewBatch()
	event := &protobuf.EventV3Dto{}
	b.pushEvent(event)

	select {
	case e := <-b.(*batch).events:
		assert.Equal(t, event, e, "Pushed event should be the same as the popped one")
	default:
		t.Error("Event was not pushed to the events channel")
	}
}

func TestPushItem(t *testing.T) {
	b := NewBatch()
	item := &protobuf.ItemDto{}
	b.pushItem(item)

	select {
	case i := <-b.(*batch).items:
		assert.Equal(t, item, i, "Pushed item should be the same as the popped one")
	default:
		t.Error("Item was not pushed to the items channel")
	}
}
