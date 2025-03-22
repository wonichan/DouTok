// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: base.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SortOrder int32

const (
	SortOrder_ASC  SortOrder = 0
	SortOrder_DESC SortOrder = 1
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	SortOrder_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

type SearchOperator int32

const (
	SearchOperator_EQ       SearchOperator = 0  // 等于
	SearchOperator_NE       SearchOperator = 1  // 不等于
	SearchOperator_GT       SearchOperator = 2  // 大于
	SearchOperator_GE       SearchOperator = 3  // 大于等于
	SearchOperator_LT       SearchOperator = 4  // 小于
	SearchOperator_LE       SearchOperator = 5  // 小于等于
	SearchOperator_LIKE     SearchOperator = 6  // 使用like的模糊匹配
	SearchOperator_WILDCARD SearchOperator = 7  // 使用通配符的模糊匹配
	SearchOperator_IN       SearchOperator = 8  // 在指定的集合中
	SearchOperator_NOT_IN   SearchOperator = 9  // 不在指定的集合中
	SearchOperator_BETWEEN  SearchOperator = 10 // 在指定的范围内
	SearchOperator_RE       SearchOperator = 11 // 正则匹配
)

// Enum value maps for SearchOperator.
var (
	SearchOperator_name = map[int32]string{
		0:  "EQ",
		1:  "NE",
		2:  "GT",
		3:  "GE",
		4:  "LT",
		5:  "LE",
		6:  "LIKE",
		7:  "WILDCARD",
		8:  "IN",
		9:  "NOT_IN",
		10: "BETWEEN",
		11: "RE",
	}
	SearchOperator_value = map[string]int32{
		"EQ":       0,
		"NE":       1,
		"GT":       2,
		"GE":       3,
		"LT":       4,
		"LE":       5,
		"LIKE":     6,
		"WILDCARD": 7,
		"IN":       8,
		"NOT_IN":   9,
		"BETWEEN":  10,
		"RE":       11,
	}
)

func (x SearchOperator) Enum() *SearchOperator {
	p := new(SearchOperator)
	*p = x
	return p
}

func (x SearchOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_base_proto_enumTypes[1].Descriptor()
}

func (SearchOperator) Type() protoreflect.EnumType {
	return &file_base_proto_enumTypes[1]
}

func (x SearchOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchOperator.Descriptor instead.
func (SearchOperator) EnumDescriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

type Metadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BizCode       int32                  `protobuf:"varint,1,opt,name=biz_code,json=bizCode,proto3" json:"biz_code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Domain        string                 `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Reason        []string               `protobuf:"bytes,4,rep,name=reason,proto3" json:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	mi := &file_base_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetBizCode() int32 {
	if x != nil {
		return x.BizCode
	}
	return 0
}

func (x *Metadata) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Metadata) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Metadata) GetReason() []string {
	if x != nil {
		return x.Reason
	}
	return nil
}

type SortField struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`                     // 用于排序的字段名称
	Order         SortOrder              `protobuf:"varint,2,opt,name=order,proto3,enum=api.SortOrder" json:"order,omitempty"` // 排序方式
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SortField) Reset() {
	*x = SortField{}
	mi := &file_base_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SortField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortField) ProtoMessage() {}

func (x *SortField) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortField.ProtoReflect.Descriptor instead.
func (*SortField) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *SortField) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SortField) GetOrder() SortOrder {
	if x != nil {
		return x.Order
	}
	return SortOrder_ASC
}

type PaginationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"` // 页码 [1, +∞)
	Size          int32                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"` // 页面大小
	Sort          []*SortField           `protobuf:"bytes,3,rep,name=sort,proto3" json:"sort,omitempty"`  // 根据字段进行排序
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_base_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PaginationRequest) GetSort() []*SortField {
	if x != nil {
		return x.Sort
	}
	return nil
}

type PaginationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`   // 当前数据的所属页码
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"` // 总页数
	Count         int32                  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"` // 总条目数
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_base_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *PaginationResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PaginationResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SearchField struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`                                // 用于搜索的字段名称
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`                                // 搜索的值
	ValueList     []string               `protobuf:"bytes,3,rep,name=value_list,json=valueList,proto3" json:"value_list,omitempty"`       // 搜索的值列表
	Operator      SearchOperator         `protobuf:"varint,4,opt,name=operator,proto3,enum=api.SearchOperator" json:"operator,omitempty"` // 操作符
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchField) Reset() {
	*x = SearchField{}
	mi := &file_base_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchField) ProtoMessage() {}

func (x *SearchField) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchField.ProtoReflect.Descriptor instead.
func (*SearchField) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{4}
}

func (x *SearchField) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SearchField) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SearchField) GetValueList() []string {
	if x != nil {
		return x.ValueList
	}
	return nil
}

func (x *SearchField) GetOperator() SearchOperator {
	if x != nil {
		return x.Operator
	}
	return SearchOperator_EQ
}

type SearchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Search        []*SearchField         `protobuf:"bytes,1,rep,name=search,proto3" json:"search,omitempty"` // 搜索条件
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_base_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{5}
}

func (x *SearchRequest) GetSearch() []*SearchField {
	if x != nil {
		return x.Search
	}
	return nil
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x22, 0x6f, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x69, 0x7a, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x62, 0x69, 0x7a, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x5f, 0x0a, 0x11, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x54, 0x0a, 0x12,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a,
	0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x39,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2a, 0x1e, 0x0a, 0x09, 0x53, 0x6f, 0x72,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x2a, 0x81, 0x01, 0x0a, 0x0e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x06, 0x0a, 0x02,
	0x45, 0x51, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02,
	0x47, 0x54, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x45, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02,
	0x4c, 0x54, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04,
	0x4c, 0x49, 0x4b, 0x45, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x49, 0x4c, 0x44, 0x43, 0x41,
	0x52, 0x44, 0x10, 0x07, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4e, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06,
	0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x45, 0x54, 0x57,
	0x45, 0x45, 0x4e, 0x10, 0x0a, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x45, 0x10, 0x0b, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x7a, 0x65, 0x6e, 0x69, 0x74, 0x68, 0x2f, 0x44, 0x6f, 0x75, 0x54, 0x6f, 0x6b, 0x2f, 0x2e,
	0x2e, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData []byte
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_base_proto_rawDesc), len(file_base_proto_rawDesc)))
	})
	return file_base_proto_rawDescData
}

var file_base_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_base_proto_goTypes = []any{
	(SortOrder)(0),             // 0: api.SortOrder
	(SearchOperator)(0),        // 1: api.SearchOperator
	(*Metadata)(nil),           // 2: api.Metadata
	(*SortField)(nil),          // 3: api.SortField
	(*PaginationRequest)(nil),  // 4: api.PaginationRequest
	(*PaginationResponse)(nil), // 5: api.PaginationResponse
	(*SearchField)(nil),        // 6: api.SearchField
	(*SearchRequest)(nil),      // 7: api.SearchRequest
}
var file_base_proto_depIdxs = []int32{
	0, // 0: api.SortField.order:type_name -> api.SortOrder
	3, // 1: api.PaginationRequest.sort:type_name -> api.SortField
	1, // 2: api.SearchField.operator:type_name -> api.SearchOperator
	6, // 3: api.SearchRequest.search:type_name -> api.SearchField
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_base_proto_rawDesc), len(file_base_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		EnumInfos:         file_base_proto_enumTypes,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
