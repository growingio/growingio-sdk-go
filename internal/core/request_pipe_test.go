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
	"testing"
	"time"

	"github.com/golang/snappy"
)

func TestPipeManager_Execute(t *testing.T) {
	txt := `The Go programming language is an open source project to make programmers more productive. 
	Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write 
	programs that get the most out of multicore and networked machines, while its novel type system 
	enables flexible and modular program construction. Go compiles quickly to machine code yet has 
	the convenience of garbage collection and the power of run-time reflection. It's a fast, statically
	 typed, compiled language that feels like a dynamically typed, interpreted language.`
	req := &Request{
		Body:      []byte(txt),
		Headers:   make(map[string]string),
		Timestamp: time.Now().UnixMilli(),
	}

	pm := getPipeManager()
	err := pm.execute(req)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	// first decrypt
	hint := byte(req.Timestamp % 256)
	decrypted := make([]byte, len(req.Body))
	for i, b := range req.Body {
		decrypted[i] = b ^ hint
	}

	// then decompress
	decompressed, err := snappy.Decode(nil, decrypted)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if !bytes.Equal(decompressed, []byte(txt)) {
		t.Fatal("Expected decrypted and decompressed body to be the same as the original, but not")
	}

	if req.Headers["X-Compress-Codec"] != "2" {
		t.Fatalf("Expected header X-Compress-Codec to be '2', but got '%s'", req.Headers["X-Compress-Codec"])
	}

	if req.Headers["X-Crypt-Codec"] != "1" {
		t.Fatalf("Expected header X-Crypt-Codec to be '1', but got '%s'", req.Headers["X-Crypt-Codec"])
	}
}
