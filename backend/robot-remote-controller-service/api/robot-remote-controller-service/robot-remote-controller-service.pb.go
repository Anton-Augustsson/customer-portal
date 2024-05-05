// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: api/robot-remote-controller-service/robot-remote-controller-service.proto

package robot_remote_controller_service

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

type SteerOperationType int32

const (
	SteerOperationType_Left  SteerOperationType = 0
	SteerOperationType_Right SteerOperationType = 1
)

// Enum value maps for SteerOperationType.
var (
	SteerOperationType_name = map[int32]string{
		0: "Left",
		1: "Right",
	}
	SteerOperationType_value = map[string]int32{
		"Left":  0,
		"Right": 1,
	}
)

func (x SteerOperationType) Enum() *SteerOperationType {
	p := new(SteerOperationType)
	*p = x
	return p
}

func (x SteerOperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SteerOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_robot_remote_controller_service_robot_remote_controller_service_proto_enumTypes[0].Descriptor()
}

func (SteerOperationType) Type() protoreflect.EnumType {
	return &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_enumTypes[0]
}

func (x SteerOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SteerOperationType.Descriptor instead.
func (SteerOperationType) EnumDescriptor() ([]byte, []int) {
	return file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescGZIP(), []int{0}
}

type SpeedOperationType int32

const (
	SpeedOperationType_Up   SpeedOperationType = 0
	SpeedOperationType_Down SpeedOperationType = 1
	SpeedOperationType_Stop SpeedOperationType = 2
)

// Enum value maps for SpeedOperationType.
var (
	SpeedOperationType_name = map[int32]string{
		0: "Up",
		1: "Down",
		2: "Stop",
	}
	SpeedOperationType_value = map[string]int32{
		"Up":   0,
		"Down": 1,
		"Stop": 2,
	}
)

func (x SpeedOperationType) Enum() *SpeedOperationType {
	p := new(SpeedOperationType)
	*p = x
	return p
}

func (x SpeedOperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpeedOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_robot_remote_controller_service_robot_remote_controller_service_proto_enumTypes[1].Descriptor()
}

func (SpeedOperationType) Type() protoreflect.EnumType {
	return &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_enumTypes[1]
}

func (x SpeedOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpeedOperationType.Descriptor instead.
func (SpeedOperationType) EnumDescriptor() ([]byte, []int) {
	return file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescGZIP(), []int{1}
}

type SteerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId int32              `protobuf:"varint,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	RobotId      int32              `protobuf:"varint,2,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
	SteerOpt     SteerOperationType `protobuf:"varint,3,opt,name=steer_opt,json=steerOpt,proto3,enum=SteerOperationType" json:"steer_opt,omitempty"`
}

func (x *SteerRequest) Reset() {
	*x = SteerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SteerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SteerRequest) ProtoMessage() {}

func (x *SteerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SteerRequest.ProtoReflect.Descriptor instead.
func (*SteerRequest) Descriptor() ([]byte, []int) {
	return file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescGZIP(), []int{0}
}

func (x *SteerRequest) GetSubscriberId() int32 {
	if x != nil {
		return x.SubscriberId
	}
	return 0
}

func (x *SteerRequest) GetRobotId() int32 {
	if x != nil {
		return x.RobotId
	}
	return 0
}

func (x *SteerRequest) GetSteerOpt() SteerOperationType {
	if x != nil {
		return x.SteerOpt
	}
	return SteerOperationType_Left
}

type SteerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId int32 `protobuf:"varint,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	RobotId      int32 `protobuf:"varint,2,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
}

func (x *SteerReply) Reset() {
	*x = SteerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SteerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SteerReply) ProtoMessage() {}

func (x *SteerReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SteerReply.ProtoReflect.Descriptor instead.
func (*SteerReply) Descriptor() ([]byte, []int) {
	return file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescGZIP(), []int{1}
}

func (x *SteerReply) GetSubscriberId() int32 {
	if x != nil {
		return x.SubscriberId
	}
	return 0
}

func (x *SteerReply) GetRobotId() int32 {
	if x != nil {
		return x.RobotId
	}
	return 0
}

type SpeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId int32              `protobuf:"varint,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	RobotId      int32              `protobuf:"varint,2,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
	SpeedOpt     SpeedOperationType `protobuf:"varint,3,opt,name=speed_opt,json=speedOpt,proto3,enum=SpeedOperationType" json:"speed_opt,omitempty"`
}

func (x *SpeedRequest) Reset() {
	*x = SpeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeedRequest) ProtoMessage() {}

func (x *SpeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeedRequest.ProtoReflect.Descriptor instead.
func (*SpeedRequest) Descriptor() ([]byte, []int) {
	return file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescGZIP(), []int{2}
}

func (x *SpeedRequest) GetSubscriberId() int32 {
	if x != nil {
		return x.SubscriberId
	}
	return 0
}

func (x *SpeedRequest) GetRobotId() int32 {
	if x != nil {
		return x.RobotId
	}
	return 0
}

func (x *SpeedRequest) GetSpeedOpt() SpeedOperationType {
	if x != nil {
		return x.SpeedOpt
	}
	return SpeedOperationType_Up
}

type SpeedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriberId int32 `protobuf:"varint,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	RobotId      int32 `protobuf:"varint,2,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
}

func (x *SpeedReply) Reset() {
	*x = SpeedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeedReply) ProtoMessage() {}

func (x *SpeedReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeedReply.ProtoReflect.Descriptor instead.
func (*SpeedReply) Descriptor() ([]byte, []int) {
	return file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescGZIP(), []int{3}
}

func (x *SpeedReply) GetSubscriberId() int32 {
	if x != nil {
		return x.SubscriberId
	}
	return 0
}

func (x *SpeedReply) GetRobotId() int32 {
	if x != nil {
		return x.RobotId
	}
	return 0
}

var File_api_robot_remote_controller_service_robot_remote_controller_service_proto protoreflect.FileDescriptor

var file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDesc = []byte{
	0x0a, 0x49, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2d, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2d, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x0c,
	0x53, 0x74, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x09,
	0x73, 0x74, 0x65, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x53, 0x74, 0x65, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73, 0x74, 0x65, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x22, 0x4c,
	0x0a, 0x0a, 0x53, 0x74, 0x65, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x80, 0x01, 0x0a,
	0x0c, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x09, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x73, 0x70, 0x65, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x22,
	0x4c, 0x0a, 0x0a, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x2a, 0x29, 0x0a,
	0x12, 0x53, 0x74, 0x65, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x65, 0x66, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x69, 0x67, 0x68, 0x74, 0x10, 0x01, 0x2a, 0x30, 0x0a, 0x12, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06,
	0x0a, 0x02, 0x55, 0x70, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x6f, 0x77, 0x6e, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x10, 0x02, 0x32, 0x6c, 0x0a, 0x1c, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x0d, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x25, 0x0a, 0x05, 0x53, 0x74, 0x65, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x53, 0x74, 0x65,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x53, 0x74, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x41, 0x75, 0x67,
	0x75, 0x73, 0x74, 0x73, 0x73, 0x6f, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescOnce sync.Once
	file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescData = file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDesc
)

func file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescGZIP() []byte {
	file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescOnce.Do(func() {
		file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescData)
	})
	return file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDescData
}

var file_api_robot_remote_controller_service_robot_remote_controller_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_robot_remote_controller_service_robot_remote_controller_service_proto_goTypes = []interface{}{
	(SteerOperationType)(0), // 0: SteerOperationType
	(SpeedOperationType)(0), // 1: SpeedOperationType
	(*SteerRequest)(nil),    // 2: SteerRequest
	(*SteerReply)(nil),      // 3: SteerReply
	(*SpeedRequest)(nil),    // 4: SpeedRequest
	(*SpeedReply)(nil),      // 5: SpeedReply
}
var file_api_robot_remote_controller_service_robot_remote_controller_service_proto_depIdxs = []int32{
	0, // 0: SteerRequest.steer_opt:type_name -> SteerOperationType
	1, // 1: SpeedRequest.speed_opt:type_name -> SpeedOperationType
	4, // 2: RobotRemoteControllerService.Speed:input_type -> SpeedRequest
	2, // 3: RobotRemoteControllerService.Steer:input_type -> SteerRequest
	5, // 4: RobotRemoteControllerService.Speed:output_type -> SpeedReply
	3, // 5: RobotRemoteControllerService.Steer:output_type -> SteerReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_robot_remote_controller_service_robot_remote_controller_service_proto_init() }
func file_api_robot_remote_controller_service_robot_remote_controller_service_proto_init() {
	if File_api_robot_remote_controller_service_robot_remote_controller_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SteerRequest); i {
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
		file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SteerReply); i {
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
		file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeedRequest); i {
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
		file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeedReply); i {
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
			RawDescriptor: file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_robot_remote_controller_service_robot_remote_controller_service_proto_goTypes,
		DependencyIndexes: file_api_robot_remote_controller_service_robot_remote_controller_service_proto_depIdxs,
		EnumInfos:         file_api_robot_remote_controller_service_robot_remote_controller_service_proto_enumTypes,
		MessageInfos:      file_api_robot_remote_controller_service_robot_remote_controller_service_proto_msgTypes,
	}.Build()
	File_api_robot_remote_controller_service_robot_remote_controller_service_proto = out.File
	file_api_robot_remote_controller_service_robot_remote_controller_service_proto_rawDesc = nil
	file_api_robot_remote_controller_service_robot_remote_controller_service_proto_goTypes = nil
	file_api_robot_remote_controller_service_robot_remote_controller_service_proto_depIdxs = nil
}
