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
//
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: event_v3.proto
// script: protoc -I=./protobuf/proto --go_out=./ ./protobuf/proto/event_v3.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_VISIT                 EventType = 0
	EventType_CUSTOM                EventType = 1
	EventType_VISITOR_ATTRIBUTES    EventType = 2
	EventType_LOGIN_USER_ATTRIBUTES EventType = 3
	EventType_CONVERSION_VARIABLES  EventType = 4
	EventType_APP_CLOSED            EventType = 5
	EventType_PAGE                  EventType = 6
	EventType_PAGE_ATTRIBUTES       EventType = 7
	EventType_VIEW_CLICK            EventType = 8
	EventType_VIEW_CHANGE           EventType = 9
	EventType_FORM_SUBMIT           EventType = 10
	EventType_ACTIVATE              EventType = 11
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0:  "VISIT",
		1:  "CUSTOM",
		2:  "VISITOR_ATTRIBUTES",
		3:  "LOGIN_USER_ATTRIBUTES",
		4:  "CONVERSION_VARIABLES",
		5:  "APP_CLOSED",
		6:  "PAGE",
		7:  "PAGE_ATTRIBUTES",
		8:  "VIEW_CLICK",
		9:  "VIEW_CHANGE",
		10: "FORM_SUBMIT",
		11: "ACTIVATE",
	}
	EventType_value = map[string]int32{
		"VISIT":                 0,
		"CUSTOM":                1,
		"VISITOR_ATTRIBUTES":    2,
		"LOGIN_USER_ATTRIBUTES": 3,
		"CONVERSION_VARIABLES":  4,
		"APP_CLOSED":            5,
		"PAGE":                  6,
		"PAGE_ATTRIBUTES":       7,
		"VIEW_CLICK":            8,
		"VIEW_CHANGE":           9,
		"FORM_SUBMIT":           10,
		"ACTIVATE":              11,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_v3_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_event_v3_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_event_v3_proto_rawDescGZIP(), []int{0}
}

type EventV3Dto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId            string            `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	UserId              string            `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionId           string            `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	DataSourceId        string            `protobuf:"bytes,5,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	EventType           EventType         `protobuf:"varint,6,opt,name=event_type,json=eventType,proto3,enum=EventType" json:"event_type,omitempty"`
	Platform            string            `protobuf:"bytes,7,opt,name=platform,proto3" json:"platform,omitempty"`
	Timestamp           int64             `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Domain              string            `protobuf:"bytes,9,opt,name=domain,proto3" json:"domain,omitempty"`
	Path                string            `protobuf:"bytes,10,opt,name=path,proto3" json:"path,omitempty"`
	Query               string            `protobuf:"bytes,11,opt,name=query,proto3" json:"query,omitempty"`
	Title               string            `protobuf:"bytes,12,opt,name=title,proto3" json:"title,omitempty"`
	ReferralPage        string            `protobuf:"bytes,13,opt,name=referral_page,json=referralPage,proto3" json:"referral_page,omitempty"`
	EventSequenceId     int32             `protobuf:"varint,15,opt,name=event_sequence_id,json=eventSequenceId,proto3" json:"event_sequence_id,omitempty"`
	ScreenHeight        int32             `protobuf:"varint,16,opt,name=screen_height,json=screenHeight,proto3" json:"screen_height,omitempty"`
	ScreenWidth         int32             `protobuf:"varint,17,opt,name=screen_width,json=screenWidth,proto3" json:"screen_width,omitempty"`
	Language            string            `protobuf:"bytes,18,opt,name=language,proto3" json:"language,omitempty"`
	SdkVersion          string            `protobuf:"bytes,19,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`
	AppVersion          string            `protobuf:"bytes,20,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	ExtraSdk            map[string]string `protobuf:"bytes,21,rep,name=extra_sdk,json=extraSdk,proto3" json:"extra_sdk,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EventName           string            `protobuf:"bytes,22,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Attributes          map[string]string `protobuf:"bytes,24,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResourceItem        *ResourceItem     `protobuf:"bytes,25,opt,name=resource_item,json=resourceItem,proto3" json:"resource_item,omitempty"`
	ProtocolType        string            `protobuf:"bytes,26,opt,name=protocol_type,json=protocolType,proto3" json:"protocol_type,omitempty"`
	TextValue           string            `protobuf:"bytes,27,opt,name=text_value,json=textValue,proto3" json:"text_value,omitempty"`
	Xpath               string            `protobuf:"bytes,28,opt,name=xpath,proto3" json:"xpath,omitempty"`
	Index               int32             `protobuf:"varint,29,opt,name=index,proto3" json:"index,omitempty"`
	Hyperlink           string            `protobuf:"bytes,30,opt,name=hyperlink,proto3" json:"hyperlink,omitempty"`
	UrlScheme           string            `protobuf:"bytes,31,opt,name=url_scheme,json=urlScheme,proto3" json:"url_scheme,omitempty"`
	AppState            string            `protobuf:"bytes,32,opt,name=app_state,json=appState,proto3" json:"app_state,omitempty"`
	NetworkState        string            `protobuf:"bytes,33,opt,name=network_state,json=networkState,proto3" json:"network_state,omitempty"`
	AppChannel          string            `protobuf:"bytes,34,opt,name=app_channel,json=appChannel,proto3" json:"app_channel,omitempty"`
	PageName            string            `protobuf:"bytes,35,opt,name=page_name,json=pageName,proto3" json:"page_name,omitempty"` // useless
	PlatformVersion     string            `protobuf:"bytes,36,opt,name=platform_version,json=platformVersion,proto3" json:"platform_version,omitempty"`
	DeviceBrand         string            `protobuf:"bytes,37,opt,name=device_brand,json=deviceBrand,proto3" json:"device_brand,omitempty"`
	DeviceModel         string            `protobuf:"bytes,38,opt,name=device_model,json=deviceModel,proto3" json:"device_model,omitempty"`
	DeviceType          string            `protobuf:"bytes,39,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	OperatingSystem     string            `protobuf:"bytes,40,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
	AppName             string            `protobuf:"bytes,42,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Latitude            float64           `protobuf:"fixed64,44,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude           float64           `protobuf:"fixed64,45,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Imei                string            `protobuf:"bytes,46,opt,name=imei,proto3" json:"imei,omitempty"`
	AndroidId           string            `protobuf:"bytes,47,opt,name=android_id,json=androidId,proto3" json:"android_id,omitempty"`
	Oaid                string            `protobuf:"bytes,48,opt,name=oaid,proto3" json:"oaid,omitempty"`
	GoogleAdvertisingId string            `protobuf:"bytes,49,opt,name=google_advertising_id,json=googleAdvertisingId,proto3" json:"google_advertising_id,omitempty"`
	Idfa                string            `protobuf:"bytes,50,opt,name=idfa,proto3" json:"idfa,omitempty"`
	Idfv                string            `protobuf:"bytes,51,opt,name=idfv,proto3" json:"idfv,omitempty"`
	Orientation         string            `protobuf:"bytes,52,opt,name=orientation,proto3" json:"orientation,omitempty"`
	ProjectKey          string            `protobuf:"bytes,53,opt,name=project_key,json=projectKey,proto3" json:"project_key,omitempty"`
	SendTime            int64             `protobuf:"varint,54,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	UserKey             string            `protobuf:"bytes,55,opt,name=user_key,json=userKey,proto3" json:"user_key,omitempty"`
	Xcontent            string            `protobuf:"bytes,56,opt,name=xcontent,proto3" json:"xcontent,omitempty"`
	TimezoneOffset      string            `protobuf:"bytes,57,opt,name=timezone_offset,json=timezoneOffset,proto3" json:"timezone_offset,omitempty"`
}

func (x *EventV3Dto) Reset() {
	*x = EventV3Dto{}
	mi := &file_event_v3_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventV3Dto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventV3Dto) ProtoMessage() {}

func (x *EventV3Dto) ProtoReflect() protoreflect.Message {
	mi := &file_event_v3_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventV3Dto.ProtoReflect.Descriptor instead.
func (*EventV3Dto) Descriptor() ([]byte, []int) {
	return file_event_v3_proto_rawDescGZIP(), []int{0}
}

func (x *EventV3Dto) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *EventV3Dto) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EventV3Dto) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *EventV3Dto) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *EventV3Dto) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_VISIT
}

func (x *EventV3Dto) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *EventV3Dto) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *EventV3Dto) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *EventV3Dto) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *EventV3Dto) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *EventV3Dto) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EventV3Dto) GetReferralPage() string {
	if x != nil {
		return x.ReferralPage
	}
	return ""
}

func (x *EventV3Dto) GetEventSequenceId() int32 {
	if x != nil {
		return x.EventSequenceId
	}
	return 0
}

func (x *EventV3Dto) GetScreenHeight() int32 {
	if x != nil {
		return x.ScreenHeight
	}
	return 0
}

func (x *EventV3Dto) GetScreenWidth() int32 {
	if x != nil {
		return x.ScreenWidth
	}
	return 0
}

func (x *EventV3Dto) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *EventV3Dto) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

func (x *EventV3Dto) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *EventV3Dto) GetExtraSdk() map[string]string {
	if x != nil {
		return x.ExtraSdk
	}
	return nil
}

func (x *EventV3Dto) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *EventV3Dto) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *EventV3Dto) GetResourceItem() *ResourceItem {
	if x != nil {
		return x.ResourceItem
	}
	return nil
}

func (x *EventV3Dto) GetProtocolType() string {
	if x != nil {
		return x.ProtocolType
	}
	return ""
}

func (x *EventV3Dto) GetTextValue() string {
	if x != nil {
		return x.TextValue
	}
	return ""
}

func (x *EventV3Dto) GetXpath() string {
	if x != nil {
		return x.Xpath
	}
	return ""
}

func (x *EventV3Dto) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *EventV3Dto) GetHyperlink() string {
	if x != nil {
		return x.Hyperlink
	}
	return ""
}

func (x *EventV3Dto) GetUrlScheme() string {
	if x != nil {
		return x.UrlScheme
	}
	return ""
}

func (x *EventV3Dto) GetAppState() string {
	if x != nil {
		return x.AppState
	}
	return ""
}

func (x *EventV3Dto) GetNetworkState() string {
	if x != nil {
		return x.NetworkState
	}
	return ""
}

func (x *EventV3Dto) GetAppChannel() string {
	if x != nil {
		return x.AppChannel
	}
	return ""
}

func (x *EventV3Dto) GetPageName() string {
	if x != nil {
		return x.PageName
	}
	return ""
}

func (x *EventV3Dto) GetPlatformVersion() string {
	if x != nil {
		return x.PlatformVersion
	}
	return ""
}

func (x *EventV3Dto) GetDeviceBrand() string {
	if x != nil {
		return x.DeviceBrand
	}
	return ""
}

func (x *EventV3Dto) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *EventV3Dto) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *EventV3Dto) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *EventV3Dto) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *EventV3Dto) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *EventV3Dto) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *EventV3Dto) GetImei() string {
	if x != nil {
		return x.Imei
	}
	return ""
}

func (x *EventV3Dto) GetAndroidId() string {
	if x != nil {
		return x.AndroidId
	}
	return ""
}

func (x *EventV3Dto) GetOaid() string {
	if x != nil {
		return x.Oaid
	}
	return ""
}

func (x *EventV3Dto) GetGoogleAdvertisingId() string {
	if x != nil {
		return x.GoogleAdvertisingId
	}
	return ""
}

func (x *EventV3Dto) GetIdfa() string {
	if x != nil {
		return x.Idfa
	}
	return ""
}

func (x *EventV3Dto) GetIdfv() string {
	if x != nil {
		return x.Idfv
	}
	return ""
}

func (x *EventV3Dto) GetOrientation() string {
	if x != nil {
		return x.Orientation
	}
	return ""
}

func (x *EventV3Dto) GetProjectKey() string {
	if x != nil {
		return x.ProjectKey
	}
	return ""
}

func (x *EventV3Dto) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *EventV3Dto) GetUserKey() string {
	if x != nil {
		return x.UserKey
	}
	return ""
}

func (x *EventV3Dto) GetXcontent() string {
	if x != nil {
		return x.Xcontent
	}
	return ""
}

func (x *EventV3Dto) GetTimezoneOffset() string {
	if x != nil {
		return x.TimezoneOffset
	}
	return ""
}

type ResourceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key        string            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Attributes map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResourceItem) Reset() {
	*x = ResourceItem{}
	mi := &file_event_v3_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceItem) ProtoMessage() {}

func (x *ResourceItem) ProtoReflect() protoreflect.Message {
	mi := &file_event_v3_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceItem.ProtoReflect.Descriptor instead.
func (*ResourceItem) Descriptor() ([]byte, []int) {
	return file_event_v3_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ResourceItem) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type EventV3List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*EventV3Dto `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *EventV3List) Reset() {
	*x = EventV3List{}
	mi := &file_event_v3_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventV3List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventV3List) ProtoMessage() {}

func (x *EventV3List) ProtoReflect() protoreflect.Message {
	mi := &file_event_v3_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventV3List.ProtoReflect.Descriptor instead.
func (*EventV3List) Descriptor() ([]byte, []int) {
	return file_event_v3_proto_rawDescGZIP(), []int{2}
}

func (x *EventV3List) GetValues() []*EventV3Dto {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_event_v3_proto protoreflect.FileDescriptor

var file_event_v3_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8a, 0x0f, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x33, 0x44, 0x74, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x73, 0x64, 0x6b, 0x18, 0x15, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x33, 0x44, 0x74, 0x6f, 0x2e, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x53, 0x64, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x53, 0x64, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x56, 0x33, 0x44, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x32, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x65, 0x78, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x78, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x78, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x72, 0x6c, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x25,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x2a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x18, 0x2e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x61, 0x69, 0x64,
	0x18, 0x30, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x61, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x15,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x31, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x64, 0x66, 0x61, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x64, 0x66, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x64, 0x66, 0x76, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x64, 0x66, 0x76, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x35, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x36, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x37, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x78, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x38, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x78, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x39, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x1a, 0x3b, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x53, 0x64, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x0e, 0x10, 0x0f,
	0x4a, 0x04, 0x08, 0x17, 0x10, 0x18, 0x4a, 0x04, 0x08, 0x29, 0x10, 0x2a, 0x4a, 0x04, 0x08, 0x2b,
	0x10, 0x2c, 0x52, 0x06, 0x67, 0x69, 0x6f, 0x5f, 0x69, 0x64, 0x52, 0x12, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x52, 0x13,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x18, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xae, 0x01,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a,
	0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x32,
	0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x33, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x33, 0x44, 0x74, 0x6f, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x2a, 0xde, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x53, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x49, 0x53, 0x49, 0x54,
	0x4f, 0x52, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x02, 0x12,
	0x19, 0x0a, 0x15, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x54,
	0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x53, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x50, 0x50, 0x5f, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x47, 0x45, 0x10, 0x06, 0x12, 0x13,
	0x0a, 0x0f, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45,
	0x53, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x43,
	0x4b, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x53, 0x55, 0x42,
	0x4d, 0x49, 0x54, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54,
	0x45, 0x10, 0x0b, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_v3_proto_rawDescOnce sync.Once
	file_event_v3_proto_rawDescData = file_event_v3_proto_rawDesc
)

func file_event_v3_proto_rawDescGZIP() []byte {
	file_event_v3_proto_rawDescOnce.Do(func() {
		file_event_v3_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_v3_proto_rawDescData)
	})
	return file_event_v3_proto_rawDescData
}

var file_event_v3_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_event_v3_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_event_v3_proto_goTypes = []any{
	(EventType)(0),       // 0: EventType
	(*EventV3Dto)(nil),   // 1: EventV3Dto
	(*ResourceItem)(nil), // 2: ResourceItem
	(*EventV3List)(nil),  // 3: EventV3List
	nil,                  // 4: EventV3Dto.ExtraSdkEntry
	nil,                  // 5: EventV3Dto.AttributesEntry
	nil,                  // 6: ResourceItem.AttributesEntry
}
var file_event_v3_proto_depIdxs = []int32{
	0, // 0: EventV3Dto.event_type:type_name -> EventType
	4, // 1: EventV3Dto.extra_sdk:type_name -> EventV3Dto.ExtraSdkEntry
	5, // 2: EventV3Dto.attributes:type_name -> EventV3Dto.AttributesEntry
	2, // 3: EventV3Dto.resource_item:type_name -> ResourceItem
	6, // 4: ResourceItem.attributes:type_name -> ResourceItem.AttributesEntry
	1, // 5: EventV3List.values:type_name -> EventV3Dto
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_event_v3_proto_init() }
func file_event_v3_proto_init() {
	if File_event_v3_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_v3_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_v3_proto_goTypes,
		DependencyIndexes: file_event_v3_proto_depIdxs,
		EnumInfos:         file_event_v3_proto_enumTypes,
		MessageInfos:      file_event_v3_proto_msgTypes,
	}.Build()
	File_event_v3_proto = out.File
	file_event_v3_proto_rawDesc = nil
	file_event_v3_proto_goTypes = nil
	file_event_v3_proto_depIdxs = nil
}
