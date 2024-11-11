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

import "github.com/golang/snappy"

type Pipe func(*request) error

type PipeManager struct {
	pipes []Pipe
}

var pipe *PipeManager

func getPipeManager() *PipeManager {
	if pipe == nil {
		pipe = &PipeManager{}
		pipe.addPipe(compress)
		pipe.addPipe(encrypt)
	}
	return pipe
}

func (pm *PipeManager) addPipe(pipe Pipe) {
	pm.pipes = append(pm.pipes, pipe)
}

func (pm *PipeManager) execute(req *request) error {
	for _, pipe := range pm.pipes {
		if err := pipe(req); err != nil {
			return err
		}
	}
	return nil
}

func compress(req *request) error {
	compressed := snappy.Encode(nil, req.Body)
	req.Body = compressed
	req.Headers["X-Compress-Codec"] = "2"
	return nil
}

func encrypt(req *request) error {
	hint := byte(req.Timestamp % 256)
	encrypted := make([]byte, len(req.Body))
	for i, b := range req.Body {
		encrypted[i] = b ^ hint
	}
	req.Body = encrypted
	req.Headers["X-Crypt-Codec"] = "1"
	return nil
}
