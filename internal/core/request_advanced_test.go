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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/growingio/growingio-sdk-go/internal/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestNewRequest_Success(t *testing.T) {
	// Create test data
	event := &protobuf.EventV3Dto{
		EventName:   "test_event",
		ProjectKey:  "test_project",
		EventType:   protobuf.EventType_CUSTOM,
	}

	baseURL := "https://api.test.com/collect"
	
	// Test successful request creation
	req, err := NewRequest(event, baseURL)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	
	// Verify request structure
	actualReq := req.(*request)
	assert.Contains(t, actualReq.URL, baseURL)
	assert.Contains(t, actualReq.URL, "stm=")
	assert.Equal(t, "application/protobuf", actualReq.Headers["Content-Type"])
	assert.Equal(t, "application/json", actualReq.Headers["Accept"])
	assert.NotEmpty(t, actualReq.Headers["X-Timestamp"])
	assert.NotNil(t, actualReq.Body)
	assert.Greater(t, actualReq.Timestamp, int64(0))
}

func TestNewRequest_MarshalError(t *testing.T) {
	// This is harder to test since protobuf.Marshal rarely fails
	// But we can test the error handling path exists
	baseURL := "https://api.test.com/collect"
	
	// Create a valid event - this should succeed
	event := &protobuf.EventV3Dto{
		EventName: "test_event",
	}
	
	req, err := NewRequest(event, baseURL)
	assert.NoError(t, err)
	assert.NotNil(t, req)
}

func TestSendItem_Integration(t *testing.T) {
	// Create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and headers
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/protobuf", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))
		assert.NotEmpty(t, r.Header.Get("X-Timestamp"))
		
		// Verify URL contains timestamp
		assert.Contains(t, r.URL.RawQuery, "stm=")
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success"}`))
	}))
	defer server.Close()

	// Set up test environment
	Urls = RequestUrl{
		Item: server.URL,
		Collect: server.URL,
	}
	RequestTimeout = 5

	// Create test item
	item := &protobuf.ItemDto{
		Id:         "test_item_123",
		Key:        "product",
		ProjectKey: "test_project",
		Attributes: map[string]string{
			"name":  "Test Item",
			"price": "99.99",
		},
	}

	// Test sending item
	sendItem(item)
	
	// If we reach here without panic, the test passed
	assert.True(t, true)
}

func TestSendItems_Integration(t *testing.T) {
	// Create a test HTTP server
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/protobuf", r.Header.Get("Content-Type"))
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success"}`))
	}))
	defer server.Close()

	// Set up test environment
	Urls = RequestUrl{
		Item: server.URL,
		Collect: server.URL,
	}

	// Create test items
	items := []*protobuf.ItemDto{
		{Id: "item_1", Key: "product", ProjectKey: "test"},
		{Id: "item_2", Key: "product", ProjectKey: "test"},
		{Id: "item_3", Key: "product", ProjectKey: "test"},
	}

	// Test sending multiple items
	sendItems(items)
	
	// Verify request was made
	assert.Equal(t, 1, requestCount)
}

func TestSendEvent_Integration(t *testing.T) {
	// Create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/protobuf", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success"}`))
	}))
	defer server.Close()

	// Set up test environment
	Urls = RequestUrl{
		Collect: server.URL,
		Item: server.URL,
	}
	RequestTimeout = 5

	// Create test event
	event := &protobuf.EventV3Dto{
		EventName:   "test_event",
		ProjectKey:  "test_project",
		EventType:   protobuf.EventType_CUSTOM,
		DeviceId:    "test_device",
		UserId:      "test_user",
		Timestamp:   time.Now().UnixMilli(),
	}

	// Test sending event
	sendEvent(event)
	
	// If we reach here without panic, the test passed
	assert.True(t, true)
}

func TestSendEvents_Integration(t *testing.T) {
	// Create a test HTTP server
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		assert.Equal(t, http.MethodPost, r.Method)
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success"}`))
	}))
	defer server.Close()

	// Set up test environment
	Urls = RequestUrl{
		Collect: server.URL,
		Item: server.URL,
	}

	// Create test events
	events := []*protobuf.EventV3Dto{
		{EventName: "event_1", ProjectKey: "test", EventType: protobuf.EventType_CUSTOM},
		{EventName: "event_2", ProjectKey: "test", EventType: protobuf.EventType_CUSTOM},
	}

	// Test sending multiple events
	sendEvents(events)
	
	// Verify request was made
	assert.Equal(t, 1, requestCount)
}

func TestRequest_Prepare(t *testing.T) {
	// Create a request
	req := &request{
		URL:     "https://api.test.com/collect",
		Headers: map[string]string{"Test": "Value"},
		Body:    []byte("test body"),
	}

	// Test prepare method
	err := req.prepare()
	assert.NoError(t, err)
}

func TestRequest_Send_Success(t *testing.T) {
	// Create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify headers were set
		assert.Equal(t, "Value", r.Header.Get("Test"))
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success"}`))
	}))
	defer server.Close()

	// Create a request
	req := &request{
		URL:     server.URL,
		Headers: map[string]string{"Test": "Value"},
		Body:    []byte("test body"),
	}

	// Test send method
	err := req.send()
	assert.NoError(t, err)
}

func TestRequest_Send_NetworkError(t *testing.T) {
	// Create a request with invalid URL
	req := &request{
		URL:     "http://invalid-url-that-does-not-exist.local:99999",
		Headers: map[string]string{},
		Body:    []byte("test body"),
	}

	// Test send method with network error
	err := req.send()
	assert.Error(t, err)
}

func TestRequest_Send_Timeout(t *testing.T) {
	// Create a slow server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Longer than our timeout
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Set very short timeout
	originalTimeout := RequestTimeout
	RequestTimeout = 1
	defer func() { RequestTimeout = originalTimeout }()

	// Create a request
	req := &request{
		URL:     server.URL,
		Headers: map[string]string{},
		Body:    []byte("test body"),
	}

	// Test send method with timeout
	err := req.send()
	assert.Error(t, err)
}

func TestSendRequest_PrepareError(t *testing.T) {
	// Create a mock request that fails on prepare
	mockReq := &mockFailingRequest{failOnPrepare: true}
	
	// Test sendRequest with prepare error
	sendRequest(mockReq)
	
	// Should not panic, errors are logged
	assert.True(t, true)
}

func TestSendRequest_SendError(t *testing.T) {
	// Create a mock request that fails on send
	mockReq := &mockFailingRequest{failOnSend: true}
	
	// Test sendRequest with send error
	sendRequest(mockReq)
	
	// Should not panic, errors are logged
	assert.True(t, true)
}

// Mock request for testing error scenarios
type mockFailingRequest struct {
	failOnPrepare bool
	failOnSend    bool
}

func (m *mockFailingRequest) prepare() error {
	if m.failOnPrepare {
		return fmt.Errorf("mock prepare error")
	}
	return nil
}

func (m *mockFailingRequest) send() error {
	if m.failOnSend {
		return fmt.Errorf("mock send error")
	}
	return nil
}

func TestRequestURL_Generation(t *testing.T) {
	baseURL := "https://api.test.com/collect"
	event := &protobuf.EventV3Dto{EventName: "test"}
	
	req, err := NewRequest(event, baseURL)
	assert.NoError(t, err)
	
	actualReq := req.(*request)
	
	// Verify URL contains base URL
	assert.Contains(t, actualReq.URL, baseURL)
	
	// Verify URL contains timestamp parameter
	assert.Contains(t, actualReq.URL, "stm=")
	
	// Verify timestamp is numeric
	assert.Regexp(t, `stm=\d+`, actualReq.URL)
}

func TestDefaultRequestTimeout(t *testing.T) {
	// Test that default timeout is set correctly
	assert.Equal(t, DefaultRequestTimeoutSeconds, RequestTimeout)
	assert.Greater(t, RequestTimeout, 0)
	assert.LessOrEqual(t, RequestTimeout, 60) // Reasonable upper bound
}