// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: proto/student.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PhoneType int32

const (
	PhoneType_CELLULAR PhoneType = 0
	PhoneType_HOME     PhoneType = 1
)

// Enum value maps for PhoneType.
var (
	PhoneType_name = map[int32]string{
		0: "CELLULAR",
		1: "HOME",
	}
	PhoneType_value = map[string]int32{
		"CELLULAR": 0,
		"HOME":     1,
	}
)

func (x PhoneType) Enum() *PhoneType {
	p := new(PhoneType)
	*p = x
	return p
}

func (x PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_student_proto_enumTypes[0].Descriptor()
}

func (PhoneType) Type() protoreflect.EnumType {
	return &file_proto_student_proto_enumTypes[0]
}

func (x PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhoneType.Descriptor instead.
func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{0}
}

type PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string    `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type   PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=student.PhoneType" json:"type,omitempty"`
}

func (x *PhoneNumber) Reset() {
	*x = PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNumber) ProtoMessage() {}

func (x *PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNumber.ProtoReflect.Descriptor instead.
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{0}
}

func (x *PhoneNumber) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *PhoneNumber) GetType() PhoneType {
	if x != nil {
		return x.Type
	}
	return PhoneType_CELLULAR
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName  string                 `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	BirthDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=birthDate,proto3" json:"birthDate,omitempty"`
	Class     string                 `protobuf:"bytes,4,opt,name=class,proto3" json:"class,omitempty"`
	Address   string                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Phones    []*PhoneNumber         `protobuf:"bytes,6,rep,name=phones,proto3" json:"phones,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{1}
}

func (x *Student) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Student) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Student) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *Student) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *Student) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Student) GetPhones() []*PhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

type UpdateStudentPhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NewPhone string    `protobuf:"bytes,2,opt,name=newPhone,proto3" json:"newPhone,omitempty"`
	Type     PhoneType `protobuf:"varint,3,opt,name=type,proto3,enum=student.PhoneType" json:"type,omitempty"`
}

func (x *UpdateStudentPhoneRequest) Reset() {
	*x = UpdateStudentPhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStudentPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentPhoneRequest) ProtoMessage() {}

func (x *UpdateStudentPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentPhoneRequest.ProtoReflect.Descriptor instead.
func (*UpdateStudentPhoneRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateStudentPhoneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStudentPhoneRequest) GetNewPhone() string {
	if x != nil {
		return x.NewPhone
	}
	return ""
}

func (x *UpdateStudentPhoneRequest) GetType() PhoneType {
	if x != nil {
		return x.Type
	}
	return PhoneType_CELLULAR
}

type UpdateStudentAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NewAddress string `protobuf:"bytes,2,opt,name=newAddress,proto3" json:"newAddress,omitempty"`
}

func (x *UpdateStudentAddressRequest) Reset() {
	*x = UpdateStudentAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStudentAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentAddressRequest) ProtoMessage() {}

func (x *UpdateStudentAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentAddressRequest.ProtoReflect.Descriptor instead.
func (*UpdateStudentAddressRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateStudentAddressRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStudentAddressRequest) GetNewAddress() string {
	if x != nil {
		return x.NewAddress
	}
	return ""
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{4}
}

func (x *SuccessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SuccessResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetStudentByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStudentByIdRequest) Reset() {
	*x = GetStudentByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentByIdRequest) ProtoMessage() {}

func (x *GetStudentByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentByIdRequest.ProtoReflect.Descriptor instead.
func (*GetStudentByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{5}
}

func (x *GetStudentByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName  string                 `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	BirthDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=birthDate,proto3" json:"birthDate,omitempty"`
	Class     string                 `protobuf:"bytes,4,opt,name=class,proto3" json:"class,omitempty"`
	Address   string                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Phones    []*PhoneNumber         `protobuf:"bytes,6,rep,name=phones,proto3" json:"phones,omitempty"`
}

func (x *CreateStudentRequest) Reset() {
	*x = CreateStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentRequest) ProtoMessage() {}

func (x *CreateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentRequest.ProtoReflect.Descriptor instead.
func (*CreateStudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{6}
}

func (x *CreateStudentRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateStudentRequest) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *CreateStudentRequest) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *CreateStudentRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateStudentRequest) GetPhones() []*PhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

type ListStudent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*Student `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *ListStudent) Reset() {
	*x = ListStudent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStudent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudent) ProtoMessage() {}

func (x *ListStudent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudent.ProtoReflect.Descriptor instead.
func (*ListStudent) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{7}
}

func (x *ListStudent) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

var File_proto_student_proto protoreflect.FileDescriptor

var file_proto_student_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0b,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x07,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a,
	0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0x6f, 0x0a, 0x19, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x1b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6e,
	0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3f, 0x0a, 0x0f, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x27, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2a,
	0x23, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x45, 0x4c, 0x4c, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f,
	0x4d, 0x45, 0x10, 0x01, 0x32, 0xf1, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_student_proto_rawDescOnce sync.Once
	file_proto_student_proto_rawDescData = file_proto_student_proto_rawDesc
)

func file_proto_student_proto_rawDescGZIP() []byte {
	file_proto_student_proto_rawDescOnce.Do(func() {
		file_proto_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_student_proto_rawDescData)
	})
	return file_proto_student_proto_rawDescData
}

var file_proto_student_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_student_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_student_proto_goTypes = []interface{}{
	(PhoneType)(0),                      // 0: student.PhoneType
	(*PhoneNumber)(nil),                 // 1: student.PhoneNumber
	(*Student)(nil),                     // 2: student.Student
	(*UpdateStudentPhoneRequest)(nil),   // 3: student.UpdateStudentPhoneRequest
	(*UpdateStudentAddressRequest)(nil), // 4: student.UpdateStudentAddressRequest
	(*SuccessResponse)(nil),             // 5: student.SuccessResponse
	(*GetStudentByIdRequest)(nil),       // 6: student.GetStudentByIdRequest
	(*CreateStudentRequest)(nil),        // 7: student.CreateStudentRequest
	(*ListStudent)(nil),                 // 8: student.ListStudent
	(*timestamppb.Timestamp)(nil),       // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),               // 10: google.protobuf.Empty
}
var file_proto_student_proto_depIdxs = []int32{
	0,  // 0: student.PhoneNumber.type:type_name -> student.PhoneType
	9,  // 1: student.Student.birthDate:type_name -> google.protobuf.Timestamp
	1,  // 2: student.Student.phones:type_name -> student.PhoneNumber
	0,  // 3: student.UpdateStudentPhoneRequest.type:type_name -> student.PhoneType
	9,  // 4: student.CreateStudentRequest.birthDate:type_name -> google.protobuf.Timestamp
	1,  // 5: student.CreateStudentRequest.phones:type_name -> student.PhoneNumber
	2,  // 6: student.ListStudent.students:type_name -> student.Student
	7,  // 7: student.StudentService.CreateStudent:input_type -> student.CreateStudentRequest
	6,  // 8: student.StudentService.GetStudentById:input_type -> student.GetStudentByIdRequest
	10, // 9: student.StudentService.GetAllStudent:input_type -> google.protobuf.Empty
	4,  // 10: student.StudentService.UpdateStudentAddress:input_type -> student.UpdateStudentAddressRequest
	3,  // 11: student.StudentService.UpdateStudentPhone:input_type -> student.UpdateStudentPhoneRequest
	2,  // 12: student.StudentService.CreateStudent:output_type -> student.Student
	2,  // 13: student.StudentService.GetStudentById:output_type -> student.Student
	8,  // 14: student.StudentService.GetAllStudent:output_type -> student.ListStudent
	2,  // 15: student.StudentService.UpdateStudentAddress:output_type -> student.Student
	2,  // 16: student.StudentService.UpdateStudentPhone:output_type -> student.Student
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_student_proto_init() }
func file_proto_student_proto_init() {
	if File_proto_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneNumber); i {
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
		file_proto_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_proto_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStudentPhoneRequest); i {
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
		file_proto_student_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStudentAddressRequest); i {
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
		file_proto_student_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_proto_student_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentByIdRequest); i {
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
		file_proto_student_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudentRequest); i {
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
		file_proto_student_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStudent); i {
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
			RawDescriptor: file_proto_student_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_student_proto_goTypes,
		DependencyIndexes: file_proto_student_proto_depIdxs,
		EnumInfos:         file_proto_student_proto_enumTypes,
		MessageInfos:      file_proto_student_proto_msgTypes,
	}.Build()
	File_proto_student_proto = out.File
	file_proto_student_proto_rawDesc = nil
	file_proto_student_proto_goTypes = nil
	file_proto_student_proto_depIdxs = nil
}
