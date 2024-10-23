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

package growingio

type CustomEventBuilder struct {
	EventName    string
	EventTime    int64
	AnonymousId  string
	LoginUserId  string
	LoginUserKey string
	Attributes   map[string]interface{}
}

type UserEventBuilder struct {
	EventTime    int64
	AnonymousId  string
	LoginUserId  string
	LoginUserKey string
	Attributes   map[string]interface{}
}

type ItemBuilder struct {
	ItemId     string
	ItemKey    string
	Attributes map[string]interface{}
}

func NewCustomEvent(eventName string) *CustomEventBuilder {
	return &CustomEventBuilder{
		EventName: eventName,
	}
}

func NewUser(loginUserId string) *UserEventBuilder {
	return &UserEventBuilder{
		LoginUserId: loginUserId,
	}
}

func NewItem(itemId string, itemKey string) *ItemBuilder {
	return &ItemBuilder{
		ItemId:  itemId,
		ItemKey: itemKey,
	}
}

func (builder *CustomEventBuilder) WithEventTime(eventTime int64) *CustomEventBuilder {
	builder.EventTime = eventTime
	return builder
}

func (builder *CustomEventBuilder) WithAnonymousId(anonymousId string) *CustomEventBuilder {
	builder.AnonymousId = anonymousId
	return builder
}

func (builder *CustomEventBuilder) WithLoginUserId(loginUserId string) *CustomEventBuilder {
	builder.LoginUserId = loginUserId
	return builder
}

func (builder *CustomEventBuilder) WithLoginUserKey(loginUserKey string) *CustomEventBuilder {
	builder.LoginUserKey = loginUserKey
	return builder
}

func (builder *CustomEventBuilder) WithAttributes(attributes map[string]interface{}) *CustomEventBuilder {
	builder.Attributes = attributes
	return builder
}

func (builder *UserEventBuilder) WithEventTime(eventTime int64) *UserEventBuilder {
	builder.EventTime = eventTime
	return builder
}

func (builder *UserEventBuilder) WithAnonymousId(anonymousId string) *UserEventBuilder {
	builder.AnonymousId = anonymousId
	return builder
}

func (builder *UserEventBuilder) WithLoginUserKey(loginUserKey string) *UserEventBuilder {
	builder.LoginUserKey = loginUserKey
	return builder
}

func (builder *UserEventBuilder) WithAttributes(attributes map[string]interface{}) *UserEventBuilder {
	builder.Attributes = attributes
	return builder
}

func (builder *ItemBuilder) WithAttributes(attributes map[string]interface{}) *ItemBuilder {
	builder.Attributes = attributes
	return builder
}
