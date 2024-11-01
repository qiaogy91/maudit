// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: apps/event/pb/model.proto

package event

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"created_at" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"created_at" gorm:"primaryKey"`
	// @gotags: validate:"required"
	Time int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty" validate:"required"`
	// @gotags: validate:"required"
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Location string `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Agent string `protobuf:"bytes,6,opt,name=agent,proto3" json:"agent,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Service string `protobuf:"bytes,7,opt,name=service,proto3" json:"service,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Resource string `protobuf:"bytes,8,opt,name=resource,proto3" json:"resource,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Action string `protobuf:"bytes,9,opt,name=action,proto3" json:"action,omitempty" validate:"required"`
	// @gotags: validate:"required"
	ResponseCode int64 `protobuf:"varint,10,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty" validate:"required"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_event_pb_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_apps_event_pb_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_apps_event_pb_model_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Event) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Event) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Event) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Event) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *Event) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Event) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Event) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Event) GetResponseCode() int64 {
	if x != nil {
		return x.ResponseCode
	}
	return 0
}

type EventSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*Event `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *EventSet) Reset() {
	*x = EventSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_event_pb_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSet) ProtoMessage() {}

func (x *EventSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_event_pb_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSet.ProtoReflect.Descriptor instead.
func (*EventSet) Descriptor() ([]byte, []int) {
	return file_apps_event_pb_model_proto_rawDescGZIP(), []int{1}
}

func (x *EventSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *EventSet) GetItems() []*Event {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_event_pb_model_proto protoreflect.FileDescriptor

var file_apps_event_pb_model_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xfc, 0x01, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x4b, 0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f, 0x6d, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_event_pb_model_proto_rawDescOnce sync.Once
	file_apps_event_pb_model_proto_rawDescData = file_apps_event_pb_model_proto_rawDesc
)

func file_apps_event_pb_model_proto_rawDescGZIP() []byte {
	file_apps_event_pb_model_proto_rawDescOnce.Do(func() {
		file_apps_event_pb_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_event_pb_model_proto_rawDescData)
	})
	return file_apps_event_pb_model_proto_rawDescData
}

var file_apps_event_pb_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_event_pb_model_proto_goTypes = []any{
	(*Event)(nil),    // 0: maudit.event.Event
	(*EventSet)(nil), // 1: maudit.event.EventSet
}
var file_apps_event_pb_model_proto_depIdxs = []int32{
	0, // 0: maudit.event.EventSet.items:type_name -> maudit.event.Event
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_event_pb_model_proto_init() }
func file_apps_event_pb_model_proto_init() {
	if File_apps_event_pb_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_event_pb_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_event_pb_model_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EventSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_event_pb_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_event_pb_model_proto_goTypes,
		DependencyIndexes: file_apps_event_pb_model_proto_depIdxs,
		MessageInfos:      file_apps_event_pb_model_proto_msgTypes,
	}.Build()
	File_apps_event_pb_model_proto = out.File
	file_apps_event_pb_model_proto_rawDesc = nil
	file_apps_event_pb_model_proto_goTypes = nil
	file_apps_event_pb_model_proto_depIdxs = nil
}
