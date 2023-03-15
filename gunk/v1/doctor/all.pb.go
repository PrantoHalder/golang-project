// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: main.go/gunk/v1/doctor/all.proto

package doctorpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// user struct
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Username  string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	IsActive  bool   `protobuf:"varint,6,opt,name=Is_active,json=IsActive,proto3" json:"IsActive,omitempty"`
	Role      string `protobuf:"bytes,7,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctor_all_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

// doctor type struct
type DoctorType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	DoctorType string `protobuf:"bytes,2,opt,name=DoctorType,proto3" json:"DoctorType,omitempty"`
}

func (x *DoctorType) Reset() {
	*x = DoctorType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorType) ProtoMessage() {}

func (x *DoctorType) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorType.ProtoReflect.Descriptor instead.
func (*DoctorType) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctor_all_proto_rawDescGZIP(), []int{1}
}

func (x *DoctorType) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DoctorType) GetDoctorType() string {
	if x != nil {
		return x.DoctorType
	}
	return ""
}

// Doctor type register request
type RegisterDoctorTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	DoctorType string `protobuf:"bytes,2,opt,name=DoctorType,proto3" json:"DoctorType,omitempty"`
}

func (x *RegisterDoctorTypeRequest) Reset() {
	*x = RegisterDoctorTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDoctorTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDoctorTypeRequest) ProtoMessage() {}

func (x *RegisterDoctorTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDoctorTypeRequest.ProtoReflect.Descriptor instead.
func (*RegisterDoctorTypeRequest) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctor_all_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterDoctorTypeRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RegisterDoctorTypeRequest) GetDoctorType() string {
	if x != nil {
		return x.DoctorType
	}
	return ""
}

// Doctor type register response
type RegisterDoctorTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *DoctorType `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *RegisterDoctorTypeResponse) Reset() {
	*x = RegisterDoctorTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDoctorTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDoctorTypeResponse) ProtoMessage() {}

func (x *RegisterDoctorTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDoctorTypeResponse.ProtoReflect.Descriptor instead.
func (*RegisterDoctorTypeResponse) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctor_all_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterDoctorTypeResponse) GetUser() *DoctorType {
	if x != nil {
		return x.User
	}
	return nil
}

// doctor login request
type DoctorLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *DoctorLoginRequest) Reset() {
	*x = DoctorLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorLoginRequest) ProtoMessage() {}

func (x *DoctorLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorLoginRequest.ProtoReflect.Descriptor instead.
func (*DoctorLoginRequest) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctor_all_proto_rawDescGZIP(), []int{4}
}

func (x *DoctorLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DoctorLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// doctor login response
type DoctorLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *DoctorLoginResponse) Reset() {
	*x = DoctorLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorLoginResponse) ProtoMessage() {}

func (x *DoctorLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctor_all_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorLoginResponse.ProtoReflect.Descriptor instead.
func (*DoctorLoginResponse) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctor_all_proto_rawDescGZIP(), []int{5}
}

func (x *DoctorLoginResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_main_go_gunk_v1_doctor_all_proto protoreflect.FileDescriptor

var file_main_go_gunk_v1_doctor_all_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x6f, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x22, 0xd5, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1d, 0x0a,
	0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x12, 0x1d, 0x0a, 0x09, 0x49, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x12, 0x18, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00,
	0x10, 0x00, 0x18, 0x00, 0x22, 0x4c, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00,
	0x18, 0x00, 0x22, 0x5b, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22,
	0x54, 0x0a, 0x1a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x58, 0x0a, 0x12, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22,
	0x47, 0x0a, 0x13, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x32, 0xd9, 0x01, 0x0a, 0x0d, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x12, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x23, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00,
	0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x70,
	0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a,
	0x03, 0x88, 0x02, 0x00, 0x42, 0x3a, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x1f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x67, 0x6f, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x3b, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x80, 0x01, 0x00, 0x88, 0x01,
	0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_go_gunk_v1_doctor_all_proto_rawDescOnce sync.Once
	file_main_go_gunk_v1_doctor_all_proto_rawDescData = file_main_go_gunk_v1_doctor_all_proto_rawDesc
)

func file_main_go_gunk_v1_doctor_all_proto_rawDescGZIP() []byte {
	file_main_go_gunk_v1_doctor_all_proto_rawDescOnce.Do(func() {
		file_main_go_gunk_v1_doctor_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_go_gunk_v1_doctor_all_proto_rawDescData)
	})
	return file_main_go_gunk_v1_doctor_all_proto_rawDescData
}

var (
	file_main_go_gunk_v1_doctor_all_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
	file_main_go_gunk_v1_doctor_all_proto_goTypes  = []interface{}{
		(*User)(nil),                       // 0: doctorpb.User
		(*DoctorType)(nil),                 // 1: doctorpb.DoctorType
		(*RegisterDoctorTypeRequest)(nil),  // 2: doctorpb.RegisterDoctorTypeRequest
		(*RegisterDoctorTypeResponse)(nil), // 3: doctorpb.RegisterDoctorTypeResponse
		(*DoctorLoginRequest)(nil),         // 4: doctorpb.DoctorLoginRequest
		(*DoctorLoginResponse)(nil),        // 5: doctorpb.DoctorLoginResponse
	}
)

var file_main_go_gunk_v1_doctor_all_proto_depIdxs = []int32{
	1, // 0: doctorpb.RegisterDoctorTypeResponse.User:type_name -> doctorpb.DoctorType
	0, // 1: doctorpb.DoctorLoginResponse.User:type_name -> doctorpb.User
	2, // 2: doctorpb.DoctorService.RegisterDoctorType:input_type -> doctorpb.RegisterDoctorTypeRequest
	4, // 3: doctorpb.DoctorService.DoctorLogin:input_type -> doctorpb.DoctorLoginRequest
	3, // 4: doctorpb.DoctorService.RegisterDoctorType:output_type -> doctorpb.RegisterDoctorTypeResponse
	5, // 5: doctorpb.DoctorService.DoctorLogin:output_type -> doctorpb.DoctorLoginResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_main_go_gunk_v1_doctor_all_proto_init() }
func file_main_go_gunk_v1_doctor_all_proto_init() {
	if File_main_go_gunk_v1_doctor_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_go_gunk_v1_doctor_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_main_go_gunk_v1_doctor_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorType); i {
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
		file_main_go_gunk_v1_doctor_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDoctorTypeRequest); i {
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
		file_main_go_gunk_v1_doctor_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDoctorTypeResponse); i {
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
		file_main_go_gunk_v1_doctor_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorLoginRequest); i {
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
		file_main_go_gunk_v1_doctor_all_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorLoginResponse); i {
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
			RawDescriptor: file_main_go_gunk_v1_doctor_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_go_gunk_v1_doctor_all_proto_goTypes,
		DependencyIndexes: file_main_go_gunk_v1_doctor_all_proto_depIdxs,
		MessageInfos:      file_main_go_gunk_v1_doctor_all_proto_msgTypes,
	}.Build()
	File_main_go_gunk_v1_doctor_all_proto = out.File
	file_main_go_gunk_v1_doctor_all_proto_rawDesc = nil
	file_main_go_gunk_v1_doctor_all_proto_goTypes = nil
	file_main_go_gunk_v1_doctor_all_proto_depIdxs = nil
}
